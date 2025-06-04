package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *shift) update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusCreated).JSON("")
	}
}

type update_shift_request struct {
}

type update_shift_response struct {
}
