package http

import (
	"context"
	jsonDelivery "emailservice/internal/delivery/http/delivery"
	"emailservice/internal/domain/delivery"
	"emailservice/internal/domain/subscriber"
	"emailservice/internal/domain/subscriber/age"
	"emailservice/internal/domain/subscriber/email"
	"emailservice/internal/domain/subscriber/name"
	"emailservice/internal/domain/subscriber/surname"
	"emailservice/pkg/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateDeliveryWithSubscribers
// @Summary Создает рассылку.
// @Description Создает рассылку.
// @Tags delivery
// @Accept  json
// @Produce json
// @Param   delivery 	body 		jsonDelivery.CrateDeliveryRequest 		    true  "Данные по рассылке"
// @Success 201			{object}    jsonDelivery.CreateDeliveryResponse				"Структура рассылки"
// @Success 200
// @Failure 400 		{object}    ErrorResponse
// @Failure 500 	    {object} 	ErrorResponse			"404 Not Found"
// @Router /delivery/ [post]
func (d *Delivery) CreateDeliveryWithSubscribers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	req := jsonDelivery.CrateDeliveryRequest{}
	if err := c.ShouldBind(&req); err != nil {
		d.logger.Info("invalid request params %s", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{Msg: "invalid request params"})
		return
	}

	templateID, err := uuid.Parse(req.TemplateID)
	if err != nil {
		d.logger.Info("invalid id from request %s", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{Msg: "invalid request params"})
		return
	}

	delivery, err := delivery.New(templateID, req.SendAt)
	if err != nil {
		d.logger.Info("error creating delivery struct %s", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{Msg: "invalid request params"})
		return
	}

	subscribers := make([]*subscriber.Subscriber, len(req.Subscribers))
	for i, subs := range req.Subscribers {
		name, err := name.NewName(subs.Name)
		if err != nil {
			d.logger.Info("invalid subscribers name: %s", err.Error())
			c.JSON(http.StatusBadRequest, ErrorResponse{Msg: fmt.Sprintf("invalid subscribers name: %s", subs.Name)})
			return
		}
		surname, err := surname.NewSurname(subs.Surname)
		if err != nil {
			d.logger.Info("invalid subscribers surname: %s", err.Error())
			c.JSON(http.StatusBadRequest, ErrorResponse{Msg: fmt.Sprintf("invalid subscribers surname: %s", subs.Surname)})
			return
		}
		email, err := email.NewEmail(subs.Email)
		if err != nil {
			d.logger.Info("invalid subscribers email: %s", err.Error())
			c.JSON(http.StatusBadRequest, ErrorResponse{Msg: fmt.Sprintf("invalid subscribers email: %s", subs.Email)})
			return
		}
		age, err := age.NewAge(subs.Age)
		if err != nil {
			d.logger.Info("invalid subscribers age: %s", err.Error())
			c.JSON(http.StatusBadRequest, ErrorResponse{Msg: fmt.Sprintf("invalid subscribers age: %d", subs.Age)})
			return
		}
		subscriber, err := subscriber.New(*name, *surname, *email, *age)
		if err != nil {
			d.logger.Info("error creating subscriber from request: %s", err.Error())
			c.JSON(http.StatusInternalServerError, ErrorResponse{Msg: "internal error"})
			return
		}
		subscribers[i] = subscriber
	}

	err = d.useCase.CreateDeliveryWithSubscribers(ctx, delivery, subscribers...)
	if err != nil {
		d.logger.Info("error creating delivery with subscribers: %s", err.Error())
		c.JSON(http.StatusInternalServerError, ErrorResponse{Msg: "internal server error"})
		return
	}

	c.JSON(http.StatusOK, jsonDelivery.CreateDeliveryResponse{ID: delivery.ID(), TemplateID: delivery.TemplateID().String(), Subject: delivery.ID().String(), CreatedAt: delivery.CreatedAt(), ModifiedAt: delivery.ModifiedAt()})
}

// MarkAsReadedBySubscriber
// @Summary Рендерит скрытую гифку для письма и помечает прочитанные письма.
// @Description Рендерит скрытую гифку для письма и помечает прочитанные письма.
// @Tags delivery
// @Accept  json
// @Produce image/gif
// @Param   deliveryId	 			path 		string 		    								true  "Идентификатор рассылки"
// @Param   subscriberId 			path 		string 		    								true  "Идентификатор подписчика"
// @Success 200
// @Router /delivery/ [get]
func (d *Delivery) MarkAsReadedBySubscriber(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancel()

	deliveryId, err := uuid.Parse(c.Param("deliveryId"))
	if err != nil {
		d.logger.Error("error parsing deliveryId: %s", err.Error())
		return
	}

	subscriberId, err := uuid.Parse(c.Param("subscriberId"))
	if err != nil {
		d.logger.Error("error parsing subscriberId: %s", err.Error())
		return
	}

	err = d.useCase.MarkAsReadedBySubscriber(ctx, deliveryId, subscriberId)
	if err != nil {
		d.logger.Error("error updating readness status: %s", err.Error())
		return
	}

	gif := helpers.CreateEmptyGif()
	c.Header("Content-Disposition", "attachment; filename=image.gif")
	c.Data(http.StatusOK, "application/octet-stream", gif)
}
