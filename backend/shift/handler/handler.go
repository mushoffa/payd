package handler

import (
	"payd/infrastructure/http/server"
	"payd/shift/domain/repository"

	"github.com/gofiber/fiber/v2"
)

type shift struct {
	http.FiberHandler
	r domain.ShiftRepository
}

func NewShiftHandler(r domain.ShiftRepository) *shift {
	return &shift{r: r}
}

func Routes(r domain.ShiftRepository) *fiber.App {
	shifts := fiber.New()

	shift := NewShiftHandler(r)
	shifts.Post("/", shift.create())
	shifts.Get("/", shift.getAll())
	shifts.Get("/:id", shift.getByID())
	shifts.Put("/:id", shift.update())
	shifts.Delete("/:id", shift.delete())

	return shifts
}
