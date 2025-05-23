package handler

import (
	"github.com/gofiber/fiber/v2"
)

func ShiftRouter(app fiber.Router) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	shifts := v1.Group("/shifts")

	shifts.Post("/", createShift())
	shifts.Get("/", getAllShifts())
	shifts.Get("/:id", getShiftByID())
	shifts.Put("/:id", updateShift())
	shifts.Delete("/:id", deleteShift())
}
