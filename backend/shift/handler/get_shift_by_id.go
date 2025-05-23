package handler

import (
	"github.com/gofiber/fiber/v2"
)

func getShiftByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusCreated).JSON("")
	}
}

type get_shift_id_response struct {
}
