package db

import "context"

type DatabaseClient interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
}
