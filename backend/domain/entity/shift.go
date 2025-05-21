package entity

import (
	"errors"
	"strings"
	"time"

	"payd/domain/valueobject"
)

const (
	MIN_SHIFT_DURATION = time.Hour * 1
	MAX_SHIFT_DURATION = time.Hour * 8
)

var (
	ErrInvalidDurationMin = errors.New("Minimum shift duration is 1 hour")
	ErrInvalidDurationMax = errors.New("Maximum shift duration is 8 hour")
	ErrInvalidRole        = errors.New("Invalid role")
)

type Shift struct {
	ID           int
	Date         valueobject.ShiftDate
	StartTime    valueobject.ShiftTime
	EndTime      valueobject.ShiftTime
	Role         string
	Location     string
	EmployeeID   int
	EmployeeName string
}

func NewShift(date valueobject.ShiftDate, s_time, e_time valueobject.ShiftTime, role, loc string) Shift {
	return Shift{
		Date:      date,
		StartTime: s_time,
		EndTime:   e_time,
		Role:      strings.ToUpper(role),
		Location:  loc,
	}
}

func (e Shift) Validate() error {

	if err := e.Date.Validate(); err != nil {
		return err
	}

	if err := e.validateTime(); err != nil {
		return err
	}

	if err := e.validateRole(); err != nil {
		return err
	}

	return nil
}

func (e Shift) validateTime() error {

	startTime := e.StartTime.T
	endTime := e.EndTime.T

	duration := endTime.Sub(startTime)

	if duration < MIN_SHIFT_DURATION {
		return ErrInvalidDurationMin
	}

	if duration > MAX_SHIFT_DURATION {
		return ErrInvalidDurationMax
	}

	return nil
}

func (e Shift) validateRole() error {
	switch e.Role {
	case valueobject.RoleBarista.String(),
		valueobject.RoleCashier.String(),
		valueobject.RoleWaitress.String():
		return nil
	default:
		return ErrInvalidRole
	}
}
