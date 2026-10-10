package route

import (
	"context"
	"time"

	"eticket-go/app/service"
	"eticket-go/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Register(app *fiber.App, eventService *service.EventService, pool *pgxpool.Pool) {
	api := app.Group("/api/v1")

	api.Get("/health", healthCheck(pool))

	api.Get("/events", eventService.List)

	api.Get("/events/:id", eventService.Get)

	api.Post("/events", eventService.Create)

	api.Put("/events/:id", eventService.Replace)

	api.Delete("/events/:id", eventService.Delete)

	api.Patch("/events/:id", eventService.Patch)
}

func healthCheck(pool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.UserContext(), 2*time.Second)

		defer cancel()

		if err := pool.Ping(ctx); err != nil {
			return helper.Fail(c, fiber.StatusServiceUnavailable, "database tidak dapat dihubungi")
		}

		return helper.Success(c, fiber.StatusOK, "server dan database berjalan", nil)
	}
}
