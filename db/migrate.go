package db

import (
	"context"
	"fmt"
	"time"

	definitions "github.com/joaomarcuslf/sucellus/definitions"
	migrations "github.com/joaomarcuslf/sucellus/migrations"
	"github.com/joaomarcuslf/sucellus/models"

	"go.mongodb.org/mongo-driver/bson"
)

func Migrate(connection definitions.DatabaseClient) {
	for _, migration := range migrations.GetList() {
		_, err := GetMigrations(connection, migration.Name)

		if err != nil {
			migration.Implementation(connection)

			SaveMigration(connection, migration.Name)
		}
	}
}

func SaveMigration(connection definitions.DatabaseClient, key string) (models.Migration, error) {
	var migration = models.Migration{
		Name: key,
		Date: definitions.Datetime{
			Time: time.Now(),
		},
	}

	collection, err := connection.Collection("migrations")

	if err != nil {
		return migration, err
	}

	_, err = collection.InsertOne(context.TODO(), migration)

	if err != nil {
		return migration, err
	}

	return migration, err
}

func GetMigrations(connection definitions.DatabaseClient, key string) (models.Migration, error) {
	var migration models.Migration

	collection, err := connection.Collection("migrations")

	if err != nil {
		return migration, fmt.Errorf("Migration not runned")
	}

	err = collection.FindOne(context.TODO(), bson.M{"name": key}).Decode(&migration)

	if err != nil {
		return migration, fmt.Errorf("Migration not runned")
	}

	return migration, err
}
