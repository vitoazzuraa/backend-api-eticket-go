package route

import (
	"eticket-go/app/service"
	"eticket-go/helper"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, eventService *service.EventService) {
	api := app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error {
		return helper.Success(c, fiber.StatusOK, "server berjalan", nil)
	})

	api.Get("/events", eventService.List)

	api.Get("/events/:id", eventService.Get)

	api.Post("/events", eventService.Create)

	api.Put("/events/:id", eventService.Replace)
}
