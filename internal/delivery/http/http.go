package http

import (
	"emailservice/internal/useCase"
	"emailservice/pkg/logger"
	"fmt"

	docs "emailservice/internal/delivery/http/swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title email service
// @version 1.0
// @description email service

// @contact.name API Support
// @contact.email kanya384@mail.ru

// @BasePath /

type Delivery struct {
	useCase useCase.UseCase
	router  *gin.Engine
	logger  *logger.Logger
}

func New(useCase useCase.UseCase, logger *logger.Logger) *Delivery {
	router := gin.New()
	return &Delivery{
		router:  router,
		useCase: useCase,
		logger:  logger,
	}
}

func (d *Delivery) GetRouter() *gin.Engine {
	return d.router
}

func (d *Delivery) Run(port int) error {
	return d.router.Run(fmt.Sprintf(":%d", port))
}

func (d *Delivery) InitRoutes() {
	d.router.POST("/template", d.CreateTemplate)
	d.router.POST("/delivery", d.CreateDeliveryWithSubscribers)
	d.router.GET("/delivery/:deliveryId/:subscriberId", d.MarkAsReadedBySubscriber)

	docs.SwaggerInfo.BasePath = "/"
	d.router.Group("/docs").Any("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
