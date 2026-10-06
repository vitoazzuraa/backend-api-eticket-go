package config

import (
	"eticket-go/app/service"
	"eticket-go/helper"
	"eticket-go/route"

	"github.com/gofiber/fiber/v2"
)

func NewApp(eventService *service.EventService) *fiber.App {
	app := fiber.New()

	route.Register(app, eventService)

	app.Use(func(c *fiber.Ctx) error {
		return helper.Fail(c, fiber.StatusNotFound, "endpoint tidak ditemukan")
	})

	return app
}
