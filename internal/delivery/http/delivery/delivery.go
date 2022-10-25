package delivery

import (
	"time"

	"github.com/google/uuid"
)

type CrateDeliveryRequest struct {
	TemplateID  string               `json:"template_id" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	Subject     string               `json:"subject" binding:"required" example:"Lorem Ipsum"`
	SendAt      time.Time            `json:"sendAt"  binding:"required" example:"2022-10-25T15:33:35.304895357+03:00"`
	Subscribers []DeliverySubscriber `json:"subscribers" binding:"required"`
}

type DeliverySubscriber struct {
	Name    string `json:"name" binding:"omitempty,required" example:"Ivan"`
	Surname string `json:"surname" binding:"omitempty,required" example:"Ivanov"`
	Email   string `json:"email" binding:"omitempty,required" example:"test01@mail.ru" format:"email"`
	Age     int    `json:"age" binding:"min=10,max=100" example:"15"`
}

type CreateDeliveryResponse struct {
	ID         uuid.UUID
	TemplateID string
	Subject    string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
