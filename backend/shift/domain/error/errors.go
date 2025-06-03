package domain

import (
	"payd/domain/error"
)

// var _ domain.Error = (*shift)(nil)

var (
	InvalidShiftDate        = "Shift date cannot be a bygone date"
	InvalidShiftTime        = "Shift time cannot be a bygone time"
	InvalidShiftDurationMin = "Minimum shift duration is 1 hour"
	InvalidShiftDurationMax = "Minimum shift duration is 1 hour"
)

var (
	ErrInvalidShiftDate        = domain.BadRequest(InvalidShiftDate)
	ErrInvalidShiftTime        = domain.BadRequest(InvalidShiftTime)
	ErrInvalidShiftDurationMin = domain.BadRequest(InvalidShiftDurationMin)
	ErrInvalidShiftDurationMax = domain.BadRequest(InvalidShiftDurationMax)
)
