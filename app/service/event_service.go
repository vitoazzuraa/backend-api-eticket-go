package service

import (
	"strconv"
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

func (s *EventService) Get(c *fiber.Ctx) error {
	rawID := c.Params("id")

	id, err := strconv.Atoi(rawID)

	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "id harus berupa angka")
	}

	for _, event := range s.events {
		if event.ID == id {
			return helper.Success(c, fiber.StatusOK, "event ditemukan", event)
		}
	}

	return helper.Fail(c, fiber.StatusNotFound, "event tidak ditemukan")
}

func (s *EventService) Create(c *fiber.Ctx) error {
	var req model.CreateEventRequest

	if err := c.BodyParser(&req); err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "body harus berupa JSON yang valid")
	}

	errs := map[string]string{}

	if req.Name == "" {
		errs["name"] = "nama wajib diisi"
	}
	if req.Venue == "" {
		errs["venue"] = "venue wajib diisi"
	}
	if req.Price <= 0 {
		errs["price"] = "harga harus lebih dari 0"
	}
	if req.Quota <= 0 {
		errs["quota"] = "kuota harus lebih dari 0"
	}

	if len(errs) > 0 {
		return helper.FailValidation(c, errs)
	}

	maxID := 0

	for _, event := range s.events {
		if event.ID > maxID {
			maxID = event.ID
		}
	}

	maxID += 1

	event := model.Event{
		ID:        maxID,
		Name:      req.Name,
		Venue:     req.Venue,
		EventDate: req.EventDate,
		Price:     req.Price,
		Quota:     req.Quota,
	}

	s.events = append(s.events, event)

	return helper.Created(c, "event berhasil dibuat", event, "/api/v1/events/"+strconv.Itoa(event.ID))
}
