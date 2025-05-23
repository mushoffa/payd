package handler

import (
	"github.com/gofiber/fiber/v2"
)

func createShift() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusCreated).JSON("")
	}
}

type create_shift_request struct {
}

type create_shift_response struct {
}
