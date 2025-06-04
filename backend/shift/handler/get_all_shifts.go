package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *shift) getAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		shifts, err := h.r.FindAll()
		if err != nil {
			return err
		}

		return h.Success(c, shifts)
	}
}
