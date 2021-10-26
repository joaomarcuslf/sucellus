package definitions

import "context"

type Cursor interface {
	Decode(val interface{}) error
	Close(ctx context.Context) error
	Next(ctx context.Context) bool
	Err() error
}

type SingleResult interface {
	Decode(v interface{}) error
	Err() error
}

type Collection interface {
	InsertOne(ctx context.Context, model interface{}) (interface{}, error)
	UpdateOne(ctx context.Context, filter interface{}, model interface{}) (interface{}, error)
	DeleteOne(ctx context.Context, model interface{}) (interface{}, error)
	FindOne(ctx context.Context, filter interface{}) SingleResult
	Find(ctx context.Context, model interface{}) (Cursor, error)
}
