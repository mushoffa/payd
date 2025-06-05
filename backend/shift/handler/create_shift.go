package handler

import (
	// "payd/domain/error"
	"payd/shift/domain/entity"
	vo "payd/shift/domain/valueobject"

	"github.com/gofiber/fiber/v2"
)

func (h *shift) create() fiber.Handler {
	request := new(create_shift_request)

	return func(c *fiber.Ctx) error {
		if err := h.ValidateBody(c, request); err != nil {
			return nil
		}

		shift := &entity.Shift{
			Date:      request.Date,
			StartTime: request.StartTime,
			EndTime:   request.EndTime,
			Role:      request.Role,
		}

		if request.Location != nil {
			shift.Location = *request.Location
		}

		if err := shift.Validate(); err != nil {
			return err
		}

		id, err := h.r.Create(c.Context(), shift)
		if err != nil {
			return err
		}

		shift.ID = id

		return h.Created(c, shift)
	}
}

type create_shift_request struct {
	Date      vo.ShiftDate `json:"date" validate:"required"`
	StartTime vo.ShiftTime `json:"start_time" validate:"required"`
	EndTime   vo.ShiftTime `json:"end_time" validate:"required"`
	Role      string       `json:"role" validate:"required"`
	Location  *string      `json:"location"`
}
