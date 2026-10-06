package service

import (
	"time"

	"eticket-go/app/model"
	"eticket-go/helper"

	"github.com/gofiber/fiber/v2"
)

type EventService struct {
	events []model.Event
}

func NewEventService() *EventService {
	return &EventService{
		events: []model.Event{
			{
				ID:        1,
				Name:      "Comifuro",
				Venue:     "Stage 1",
				EventDate: time.Date(2026, 9, 1, 19, 0, 0, 0, time.UTC),
				Price:     20000,
				Quota:     100,
			},
			{
				ID:        2,
				Name:      "Festival Musik",
				Venue:     "Stage 2",
				EventDate: time.Date(2026, 9, 1, 20, 0, 0, 0, time.UTC),
				Price:     10000,
				Quota:     100,
			},
		},
	}
}

func (s *EventService) List(c *fiber.Ctx) error {
	return helper.Success(c, fiber.StatusOK, "daftar event berhasil diambil", s.events)
}
