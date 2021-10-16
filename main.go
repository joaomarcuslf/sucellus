package main

import (
	"context"

	configs "github.com/joaomarcuslf/sucellus/configs"
	db "github.com/joaomarcuslf/sucellus/db"
	"github.com/joaomarcuslf/sucellus/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	c := configs.GetConfig()

	connection := db.NewMongoConnection(c.Database)

	server := server.NewServer(c.Port, connection)

	ctx := context.Background()

	connection.Connect(ctx)

	defer connection.Close(ctx)

	go db.Migrate(ctx, connection)

	server.Run()
}
