package handler

import (
	"github.com/gofiber/fiber/v2"
)

func getAllShifts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusCreated).JSON("")
	}
}

type get_shifts_request struct {
}

type get_shifts_response struct {
}
