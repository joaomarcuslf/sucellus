package definitions

import (
	"context"
	"io"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Query(ctx context.Context, filter bson.M) ([]interface{}, error)
	Insert(ctx context.Context, model interface{}) error
	Get(ctx context.Context, id string) (interface{}, error)
	Set(ctx context.Context, uid primitive.ObjectID, model interface{}) error
	Create(ctx context.Context, body io.Reader) (interface{}, error)
	Update(ctx context.Context, id string, body io.Reader) error
	Delete(ctx context.Context, id string) error
	Validate(ctx context.Context, model interface{}) error
}
