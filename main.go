package main

import (
	"context"
	"log"

	"eticket-go/app/service"
	"eticket-go/config"
	"eticket-go/database"
)

func main() {
	config.LoadEnv()

	pool, err := database.NewPool(context.Background())

	if err != nil {
		log.Fatalf("database: %v", err)
	}

	defer pool.Close()

	eventService := service.NewEventService()

	app := config.NewApp(eventService, pool)

	log.Fatal(app.Listen(":" + config.GetEnv("APP_PORT", "3000")))
}
