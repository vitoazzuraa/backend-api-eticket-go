package service

import (
	"errors"
	"strconv"

	"eticket-go/app/model"
	"eticket-go/app/repository"
	"eticket-go/helper"

	"github.com/gofiber/fiber/v2"
)

type EventService struct {
	repo repository.EventRepository
}

func NewEventService(repo repository.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) List(c *fiber.Ctx) error {
	ctx, cancel := helper.RequestContext(c)

	defer cancel()

	events, err := s.repo.FindAll(ctx)

	if err != nil {
		return helper.Fail(c, fiber.StatusInternalServerError, "gagal mengambil data event")
	}

	return helper.Success(c, fiber.StatusOK, "daftar event berhasil diambil", events)
}

func (s *EventService) Get(c *fiber.Ctx) error {
	ctx, cancel := helper.RequestContext(c)

	defer cancel()

	rawID := c.Params("id")

	id, err := strconv.Atoi(rawID)

	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "id harus berupa angka")
	}

	event, err := s.repo.FindByID(ctx, id)

	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return helper.Fail(c, fiber.StatusNotFound, "event tidak ditemukan")
		}

		return helper.Fail(c, fiber.StatusInternalServerError, "gagal mengambil data event")
	}

	return helper.Success(c, fiber.StatusOK, "event ditemukan", event)
}

func (s *EventService) Create(c *fiber.Ctx) error {
	ctx, cancel := helper.RequestContext(c)

	defer cancel()

	var req model.CreateEventRequest

	if err := c.BodyParser(&req); err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "body harus berupa JSON yang valid")
	}

	if errs := ValidateEvent(req.Name, req.Venue, req.Price, req.Quota); len(errs) > 0 {
		return helper.FailValidation(c, errs)
	}

	event := model.Event{
		Name:      req.Name,
		Venue:     req.Venue,
		EventDate: req.EventDate,
		Price:     req.Price,
		Quota:     req.Quota,
	}

	newEvent, err := s.repo.Create(ctx, event)

	if err != nil {
		return helper.Fail(c, fiber.StatusInternalServerError, "gagal menyimpan event")
	}

	return helper.Created(c, "event berhasil dibuat", newEvent, "/api/v1/events/"+strconv.Itoa(newEvent.ID))
}

func (s *EventService) Replace(c *fiber.Ctx) error {
	ctx, cancel := helper.RequestContext(c)

	defer cancel()

	rawID := c.Params("id")

	id, err := strconv.Atoi(rawID)

	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "id harus berupa angka")
	}

	var req model.ReplaceEventRequest

	if err := c.BodyParser(&req); err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "body harus berupa JSON yang valid")
	}

	if errs := ValidateEvent(req.Name, req.Venue, req.Price, req.Quota); len(errs) > 0 {
		return helper.FailValidation(c, errs)
	}

	event := model.Event{
		ID:        id,
		Name:      req.Name,
		Venue:     req.Venue,
		EventDate: req.EventDate,
		Price:     req.Price,
		Quota:     req.Quota,
	}

	newEvent, err := s.repo.Update(ctx, event)

	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return helper.Fail(c, fiber.StatusNotFound, "event tidak ditemukan")
		}

		return helper.Fail(c, fiber.StatusInternalServerError, "gagal menyimpan event")
	}

	return helper.Success(c, fiber.StatusOK, "event berhasil diubah", newEvent)
}

func (s *EventService) Delete(c *fiber.Ctx) error {
	ctx, cancel := helper.RequestContext(c)

	defer cancel()

	rawID := c.Params("id")

	id, err := strconv.Atoi(rawID)

	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "id harus berupa angka")
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return helper.Fail(c, fiber.StatusNotFound, "event tidak ditemukan")
		}

		return helper.Fail(c, fiber.StatusInternalServerError, "gagal menghapus event")
	}

	return helper.NoContent(c)
}

func (s *EventService) Patch(c *fiber.Ctx) error {
	ctx, cancel := helper.RequestContext(c)

	defer cancel()

	rawID := c.Params("id")

	id, err := strconv.Atoi(rawID)

	if err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "id harus berupa angka")
	}

	var req model.PatchEventRequest

	if err := c.BodyParser(&req); err != nil {
		return helper.Fail(c, fiber.StatusBadRequest, "body harus berupa JSON yang valid")
	}

	if req.Name == nil && req.Venue == nil && req.EventDate == nil && req.Price == nil && req.Quota == nil {
		return helper.Fail(c, fiber.StatusBadRequest, "tidak ada field yang diubah")
	}

	event, err := s.repo.FindByID(ctx, id)

	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return helper.Fail(c, fiber.StatusNotFound, "event tidak ditemukan")
		}

		return helper.Fail(c, fiber.StatusInternalServerError, "gagal mengambil event")
	}

	if req.Name != nil {
		event.Name = *req.Name
	}
	if req.Venue != nil {
		event.Venue = *req.Venue
	}
	if req.EventDate != nil {
		event.EventDate = *req.EventDate
	}
	if req.Price != nil {
		event.Price = *req.Price
	}
	if req.Quota != nil {
		event.Quota = *req.Quota
	}

	newEvent, err := s.repo.Update(ctx, event)

	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return helper.Fail(c, fiber.StatusNotFound, "event tidak ditemukan")
		}

		return helper.Fail(c, fiber.StatusInternalServerError, "gagal menyimpan event")
	}

	return helper.Success(c, fiber.StatusOK, "event berhasil diubah", newEvent)
}
