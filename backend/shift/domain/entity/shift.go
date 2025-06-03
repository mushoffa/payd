package entity

import (
	"strings"
	"time"

	"payd/domain/error"
	shift "payd/shift/domain/error"
	vo "payd/shift/domain/valueobject"
)

const (
	MIN_SHIFT_DURATION = time.Hour * 1
	MAX_SHIFT_DURATION = time.Hour * 8
)

type Shift struct {
	ID           int          `json:"id,omitempty"`
	Date         vo.ShiftDate `json:"date"`
	StartTime    vo.ShiftTime `json:"start_time"`
	EndTime      vo.ShiftTime `json:"end_time"`
	Role         string       `json:"role"`
	Location     string       `json:"location,omitempty"`
	EmployeeID   int          `json:"employee_id,omitempty"`
	EmployeeName string       `json:"employee_name,omitempty"`
}

func NewShift(date, start, end, role, location string) (*Shift, domain.Error) {
	_date, err := vo.NewShiftDate(date)
	if err != nil {
		return nil, err
	}

	_start, err := vo.NewShiftTime(start)
	if err != nil {
		return nil, err
	}

	_end, err := vo.NewShiftTime(end)
	if err != nil {
		return nil, err
	}

	shift := &Shift{
		Date:      _date,
		StartTime: _start,
		EndTime:   _end,
		Role:      strings.ToUpper(role),
		Location:  strings.ToUpper(location),
	}

	return shift, nil
}

func (e *Shift) Validate() domain.Error {
	if err := e.validateDuration(); err != nil {
		return err
	}

	return nil
}

func (e *Shift) validateDuration() domain.Error {
	start := e.StartTime.Time()
	end := e.EndTime.Time()

	duration := end.Sub(start)

	if duration < MIN_SHIFT_DURATION {
		return shift.ErrInvalidShiftDurationMin
	}

	if duration > MAX_SHIFT_DURATION {
		return shift.ErrInvalidShiftDurationMax
	}

	return nil
}
