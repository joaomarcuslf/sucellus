package db

import (
	"context"
	"fmt"

	configs "github.com/joaomarcuslf/sucellus/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	client        *mongo.Client
	clientOptions *options.ClientOptions
}

func NewMongoConnection(config configs.DatabaseConfig) DatabaseClient {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s?authSource=admin&ssl=false&&authMechanism=SCRAM-SHA-256",
		config.Username,
		config.Password,
		config.Url,
		config.Port,
		config.Database,
	)

	clientOptions := options.Client().ApplyURI(uri)

	return &MongoConnection{
		clientOptions: clientOptions,
	}
}

func (c *MongoConnection) Collection(database, collection string) (*mongo.Collection, error) {
	return c.client.Database(database).Collection(collection), nil
}
func (c *MongoConnection) Connect(ctx context.Context) error {
	client, err := mongo.Connect(ctx, c.clientOptions)

	if err != nil {
		return err
	}

	c.client = client

	fmt.Println("Connected to MongoDB!")

	return nil
}

func (c *MongoConnection) Close(ctx context.Context) error {
	fmt.Println("Disconnecting from MongoDB!")
	return c.client.Disconnect(ctx)
}
