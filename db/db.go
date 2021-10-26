package db

import (
	"os"

	configs "github.com/joaomarcuslf/sucellus/configs"
	definitions "github.com/joaomarcuslf/sucellus/definitions"
)

func NewDbConnection(config configs.DatabaseConfig) definitions.DatabaseClient {
	if os.Getenv("DB_DRIVER") == "mongo" {
		return NewMongoConnection(config)
	}

	return NewMongoConnection(config)
}
