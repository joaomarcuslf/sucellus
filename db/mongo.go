package db

import (
	"context"
	"fmt"

	configs "github.com/joaomarcuslf/sucellus/configs"
	definitions "github.com/joaomarcuslf/sucellus/definitions"
	errors "github.com/joaomarcuslf/sucellus/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	client        *mongo.Client
	clientOptions *options.ClientOptions
	database      string
}

func NewMongoConnection(config configs.DatabaseConfig) definitions.DatabaseClient {
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
		database:      config.Database,
	}
}

func (c *MongoConnection) Collection(collection string) (*mongo.Collection, error) {
	return c.client.Database(c.database).Collection(collection), nil
}
func (c *MongoConnection) Connect(ctx context.Context) error {
	client, err := mongo.Connect(ctx, c.clientOptions)

	if err != nil {
		return errors.FormatError("MONGO_ERROR", "Connection refused", err)
	}

	c.client = client

	fmt.Println("Connected to MongoDB!")

	return nil
}

func (c *MongoConnection) Close(ctx context.Context) error {
	fmt.Println("Disconnecting from MongoDB!")
	err := c.client.Disconnect(ctx)

	if err != nil {
		return errors.FormatError("MONGO_ERROR", "Error disconnecting from MongoDB", err)
	}

	return nil
}
