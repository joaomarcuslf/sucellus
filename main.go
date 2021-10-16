package main

import (
	"context"

	configs "github.com/joaomarcuslf/sucellus/configs"
	db "github.com/joaomarcuslf/sucellus/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	c := configs.GetConfig()

	mongo := db.NewMongoConnection(c.Database)

	ctx := context.Background()

	mongo.Connect(ctx)

	defer mongo.Close(ctx)
}
