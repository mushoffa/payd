package handler

import (
	"github.com/gofiber/fiber/v2"
)

func deleteShift() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusCreated).JSON("")
	}
}

type delete_shift_request struct {
}

type delete_shift_response struct {
}
