package main

import (
	"log"

	"eticket-go/config"
)

func main() {
	app := config.NewApp()

	log.Fatal(app.Listen(":3000"))
}
