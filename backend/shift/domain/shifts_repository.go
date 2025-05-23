package domain

import (
	"payd/domain/entity"
)

type ShiftsRepository interface {
	Save(*entity.Shift) (int, error)
	GetAll() ([]entity.Shift, error)
	GetByID(int) (entity.Shift, error)
	Update(*entity.Shift) error
	Delete(int) error
}
