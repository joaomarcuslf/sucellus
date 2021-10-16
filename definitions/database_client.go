package definitions

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseClient interface {
	Collection(collection string) (*mongo.Collection, error)
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
}
