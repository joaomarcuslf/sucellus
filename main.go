package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	configs "github.com/joaomarcuslf/sucellus/configs"
	db "github.com/joaomarcuslf/sucellus/db"
	"github.com/joaomarcuslf/sucellus/run"
	"github.com/joaomarcuslf/sucellus/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	config := configs.GetConfig()

	connection := db.NewMongoConnection(config.Database)

	server := server.NewServer(config.Port, connection)

	ctx, _ := context.WithCancel(context.Background())

	connection.Connect(ctx)

	go db.Migrate(ctx, connection)
	go run.StartServices(ctx, connection)

	go func() {
		oscall := <-c
		log.Printf("Stopping server from action: %+v", oscall)

		run.StopServices(ctx, connection)
		connection.Close(ctx)

		os.Exit(0)
	}()

	server.Run()
}
