package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *shift) delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return h.BadRequest(c, "Invalid id format, must be numeric")
		}

		if err := h.r.Delete(c.Context(), id); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		return h.Success(c, "")
	}
}
