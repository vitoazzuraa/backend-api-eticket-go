package main

import (
	"log"

	"eticket-go/app/service"
	"eticket-go/config"
)

func main() {
	eventService := service.NewEventService()

	app := config.NewApp(eventService)

	log.Fatal(app.Listen(":3000"))
}
