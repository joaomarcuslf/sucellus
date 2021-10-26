package definitions

import (
	"context"
)

type DatabaseClient interface {
	Collection(collection string) (Collection, error)
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
}
