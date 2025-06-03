package http

import (
	"github.com/gofiber/fiber/v2"
)

type response struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type FiberHandler struct {
}

func (h *FiberHandler) SuccessResponse(data interface{}) response {
	r := response{
		Code:    "00",
		Message: "Success",
	}

	if (data != nil) && (data != "") {
		r.Data = data
	}

	return r
}

func (h *FiberHandler) Success(c *fiber.Ctx, data interface{}) error {
	r := h.SuccessResponse(data)
	return c.Status(fiber.StatusOK).JSON(r)
}
