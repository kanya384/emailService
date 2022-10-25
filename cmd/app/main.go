package main

import (
	"context"
	"emailservice/internal/config"
	delivery "emailservice/internal/delivery/http"
	repository "emailservice/internal/repository/postgres"
	"emailservice/internal/useCase"
	"emailservice/pkg/helpers"
	lg "emailservice/pkg/logger"
	"emailservice/pkg/mail"
	"emailservice/pkg/psql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	serviceName    = "emailservice"
	logsPath       = "vars/logs"
	workerInterval = time.Duration(time.Second * 5)
)

func main() {
	cfg, err := config.InitConfig("")
	if err != nil {
		panic(err)
	}

	err = helpers.CreatePathIfNotExists(logsPath)
	if err != nil {
		panic(err)
	}

	err = helpers.CreatePathIfNotExists(cfg.FileStore.Path)
	if err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(fmt.Sprintf("%s/%s", logsPath, "logs.log"), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logger := lg.New(cfg.Log.Level, serviceName, logFile)

	dsn := helpers.PostgresConnectionString(cfg.Pg.User, cfg.Pg.Pass, cfg.Pg.Host, cfg.Pg.Port, cfg.Pg.DbName)
	err = helpers.MigrationsUP(dsn, "file://migrations")
	if err != nil {
		logger.Fatalf("migrations error: %w", err)
	}

	pg, err := psql.New(dsn, psql.MaxPoolSize(cfg.Pg.PoolMax))
	if err != nil {
		logger.Fatalf("postgres connection error: %w", err)
	}

	repository, err := repository.NewRepository(pg)
	if err != nil {
		logger.Fatalf("error init repository: %w", err)
	}
	sender := mail.NewSender(cfg.Email.Host, cfg.Email.Port, cfg.Email.Login, cfg.Email.Pass)

	useCase := useCase.New(*repository, logger, sender, cfg.FileStore.Path, cfg.Email.Login, cfg.App.Host)

	done := make(chan struct{})
	defer close(done)

	go useCase.ProcessDelivery(context.Background(), done, workerInterval)

	delivery := delivery.New(useCase, logger)
	delivery.InitRoutes()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: delivery.GetRouter(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	done <- struct{}{}
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
