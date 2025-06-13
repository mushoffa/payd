package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *shift) getByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		childCtx, span := h.Trace(c.UserContext(), "Controller.GetByID")
		defer span.End()

		id, err := c.ParamsInt("id")
		if err != nil {
			return h.BadRequest(c, "Invalid id format, must be numeric")
		}

		shift, err := h.r.FindByID(childCtx, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Data not found"})
		}

		return h.Success(c, shift)
	}
}
