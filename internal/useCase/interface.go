package useCase

import (
	"context"
	"emailservice/internal/domain/delivery"
	"emailservice/internal/domain/subscriber"
	"emailservice/internal/domain/template"

	"github.com/google/uuid"
)

type UseCase interface {
	CreateTemplate(ctx context.Context, templateContent []byte) (template *template.Template, err error)

	CreateDeliveryWithSubscribers(ctx context.Context, delivery *delivery.Delivery, subscribers ...*subscriber.Subscriber) (err error)
	MarkAsReadedBySubscriber(ctx context.Context, deliveryID uuid.UUID, subscriberID uuid.UUID) (err error)
}
