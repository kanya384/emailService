package helpers

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func PostgresConnectionString(user, pass, host, port, dbName string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres",
		user,
		pass,
		host,
		port,
		dbName)
}

func MigrationsUP(dsn, path string) (err error) {
	m, err := migrate.New(path, dsn)
	if err != nil {
		return err
	}
	m.Up()
	return nil
}

func CreatePathIfNotExists(path string) (err error) {
	if _, err := os.Stat(path); err != nil {
		err = os.MkdirAll(path, os.ModePerm)
		return err
	}
	return
}

func CreateEmptyGif() []byte {
	var buf bytes.Buffer
	gif.Encode(&buf, image.Rect(0, 0, 1, 1), nil)
	return buf.Bytes()
}
