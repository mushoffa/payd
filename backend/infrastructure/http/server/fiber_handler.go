package http

import (
	"payd/infrastructure/http/server/middleware"
	"payd/infrastructure/trace/embedded"

	"github.com/gofiber/fiber/v2"
)

var validator = middleware.NewValidator()

type response struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type FiberHandler struct {
	embedded.Monitor
}

func (h *FiberHandler) ValidateBody(c *fiber.Ctx, body interface{}) error {
	if err := c.BodyParser(body); err != nil {
		return err
	}

	if err := validator.Struct(body); err != nil {
		h.BadRequest(c, err.Error())
		return err
	}

	return nil
}

func (h *FiberHandler) Success(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(h.SuccessResponse(data))
}

func (h *FiberHandler) Created(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(h.SuccessResponse(data))
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

func (h *FiberHandler) BadRequest(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": data})
}
