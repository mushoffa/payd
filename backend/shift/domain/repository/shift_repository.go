package domain

import (
	"context"

	"payd/shift/domain/entity"
)

type ShiftRepository interface {
	Create(context.Context, *entity.Shift) (int, error)
	FindAll(context.Context) ([]entity.Shift, error)
	FindByID(context.Context, int) (entity.Shift, error)
	Update(context.Context, *entity.Shift) error
	Delete(context.Context, int) error
}
