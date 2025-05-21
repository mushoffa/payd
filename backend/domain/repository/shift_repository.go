package domain

import (
	"payd/domain/entity"
	"time"
)

type ShiftRepository interface {
	Save(*entity.Shift) (int, error)
	FindByID(int) (entity.Shift, error)
	FindByTodayWithID(int) (entity.Shift, error)
	FindByDateRangeWithID(int, time.Time, time.Time) ([]entity.Shift, error)
}
