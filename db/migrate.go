package db

import (
	"context"
	"time"

	definitions "github.com/joaomarcuslf/sucellus/definitions"
	errors "github.com/joaomarcuslf/sucellus/errors"
	migrations "github.com/joaomarcuslf/sucellus/migrations"
	"github.com/joaomarcuslf/sucellus/models"

	"go.mongodb.org/mongo-driver/bson"
)

func Migrate(ctx context.Context, connection definitions.DatabaseClient) {
	for _, migration := range migrations.GetList() {
		_, err := GetMigrations(ctx, connection, migration.Name)

		if err != nil {
			migration.Implementation(connection)

			SaveMigration(ctx, connection, migration.Name)
		}
	}
}

func SaveMigration(ctx context.Context, connection definitions.DatabaseClient, key string) (models.Migration, error) {
	var migration = models.Migration{
		Name: key,
		Date: definitions.Datetime{
			Time: time.Now(),
		},
	}

	collection, err := connection.Collection("migrations")

	if err != nil {
		return migration, errors.FormatError("MIGRATION_ERROR", "Connection refused", err)
	}

	_, err = collection.InsertOne(ctx, migration)

	if err != nil {
		return migration, errors.FormatError("MIGRATION_ERROR", "Could not insert data", err)
	}

	return migration, err
}

func GetMigrations(ctx context.Context, connection definitions.DatabaseClient, key string) (models.Migration, error) {
	var migration models.Migration

	collection, err := connection.Collection("migrations")

	if err != nil {
		return migration, errors.FormatError("MIGRATION_ERROR", "Connection refused", err)
	}

	err = collection.FindOne(ctx, bson.M{"name": key}).Decode(&migration)

	if err != nil {
		return migration, errors.FormatError("MIGRATION_ERROR", "Migration already exists", err)
	}

	return migration, nil
}
