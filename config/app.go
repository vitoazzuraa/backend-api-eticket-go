package config

import (
	"eticket-go/app/service"
	"eticket-go/helper"
	"eticket-go/route"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewApp(eventService *service.EventService, pool *pgxpool.Pool) *fiber.App {
	app := fiber.New()

	route.Register(app, eventService, pool)

	app.Use(func(c *fiber.Ctx) error {
		return helper.Fail(c, fiber.StatusNotFound, "endpoint tidak ditemukan")
	})

	return app
}
