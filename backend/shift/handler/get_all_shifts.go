package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (h *shift) getAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		childCtx, span := h.Trace(c.UserContext(), "Controller.GetAll")
		defer span.End()

		shifts, err := h.r.FindAll(childCtx)
		if err != nil {
			return err
		}

		return h.Success(c, shifts)
	}
}
