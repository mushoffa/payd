package handler

import (
	"payd/infrastructure/http/server"
	"payd/shift/domain/repository"

	"github.com/gofiber/fiber/v2"
)

const (
	v1_api = "/api/v1/shifts"
	v2_api = "/api/v2/shifts"
)

type ShiftHandler interface {
	RegisterV1() (string, *fiber.App)
	RegisterV2() (string, *fiber.App)
}

type shift struct {
	http.FiberHandler
	r domain.ShiftRepository
}

func NewShiftHandler(r domain.ShiftRepository) ShiftHandler {
	h := &shift{r: r}
	h.Init()
	return h
}

func (h *shift) RegisterV1() (string, *fiber.App) {
	shifts := fiber.New()
	shifts.Post("/", h.create())
	shifts.Get("/", h.getAll())
	shifts.Get("/:id", h.getByID())
	shifts.Put("/:id", h.update())
	shifts.Delete("/:id", h.delete())

	return v1_api, shifts
}

func (h *shift) RegisterV2() (string, *fiber.App) {
	shifts := fiber.New()
	shifts.Post("/", h.create())
	shifts.Get("/", h.getAll())
	shifts.Get("/:id", h.getByID())
	shifts.Put("/:id", h.update())
	shifts.Delete("/:id", h.delete())

	return v2_api, shifts
}
