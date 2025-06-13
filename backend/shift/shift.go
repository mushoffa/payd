package shift

import (
	"payd/infrastructure/database"
	"payd/shift/handler"
	"payd/shift/infrastructure/repository"
)

type Shift struct {
	handlres handler.ShiftHandler
}

func New(database database.DatabaseService) *Shift {
	r := repository.NewShiftsRepository(database)
	h := handler.NewShiftHandler(r)
	return &Shift{h}
}

func (s *Shift) Handler() handler.ShiftHandler {
	return s.handlres
}
