package repository

import (
	"context"
	"emailservice/internal/domain/delivery"
	"emailservice/internal/domain/subscriber"
	"emailservice/internal/domain/template"

	"github.com/google/uuid"
)

type Delivery interface {
	CreateDelivery(ctx context.Context, delivery *delivery.Delivery) (err error)
	CreateDeliveryWithSubscribers(ctx context.Context, delivery *delivery.Delivery, subscribers ...*subscriber.Subscriber) (err error)
	UpdateDelivery(ctx context.Context, ID uuid.UUID, updateFn func(delivery *delivery.Delivery) (*delivery.Delivery, error)) (delivery *delivery.Delivery, err error)
	DeleteDelivery(ctx context.Context, ID uuid.UUID) (err error)
	MarkAsReadedBySubscriber(ctx context.Context, deliveryID uuid.UUID, subscriberID uuid.UUID) (err error)

	ReadDeliveryById(ctx context.Context, ID uuid.UUID) (delivery *delivery.Delivery, err error)
	ReadDeliveryTasks(ctx context.Context) (list []*delivery.Delivery, err error)
	ReadDeliverySubscribers(ctx context.Context, deliveryID uuid.UUID) (list []*subscriber.Subscriber, err error)
}

type Subscriber interface {
	CreateSubscriber(ctx context.Context, subscriber *subscriber.Subscriber) (err error)
	UpdateSubscriber(ctx context.Context, ID uuid.UUID, updateFn func(subscriber *subscriber.Subscriber) (*subscriber.Subscriber, error)) (subscriber *subscriber.Subscriber, err error)
	DeleteSubscriber(ctx context.Context, ID uuid.UUID) (err error)
	ReadeSubscriberById(ctx context.Context, ID uuid.UUID) (subscriber *subscriber.Subscriber, err error)
}

type Template interface {
	CreateTemplate(ctx context.Context, template *template.Template) (err error)
	UpdateTemplate(ctx context.Context, ID uuid.UUID, updateFn func(template *template.Template) (*template.Template, error)) (template *template.Template, err error)
	DeleteTemplate(ctx context.Context, ID uuid.UUID) (err error)
	ReadTemplateById(ctx context.Context, ID uuid.UUID) (template *template.Template, err error)
}
