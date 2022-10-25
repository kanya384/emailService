package subscriber

import (
	"context"
	"emailservice/internal/domain/subscriber"

	"github.com/jackc/pgx/v5"
)

type Subscriber interface {
	CreateSubscriberTx(ctx context.Context, tx pgx.Tx, subscriber ...*subscriber.Subscriber) ([]*subscriber.Subscriber, error)
}
