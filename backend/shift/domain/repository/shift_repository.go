package domain

import (
	"payd/shift/domain/entity"
)

type ShiftRepository interface {
	Create(*entity.Shift) (int, error)
	FindAll() ([]entity.Shift, error)
	FindByID(int) (entity.Shift, error)
	Update(*entity.Shift) error
	Delete(int) error
}
