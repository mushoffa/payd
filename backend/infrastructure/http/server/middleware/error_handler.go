package middleware

import (
	"payd/domain/error"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx) error {
	err := c.Next()

	if err != nil {

		if e, ok := err.(domain.Error); ok {
			return c.Status(e.Status()).JSON(fiber.Map{
				"message": e.Message(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return nil
}
