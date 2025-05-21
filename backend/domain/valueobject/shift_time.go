package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidTime       = errors.New("Invalid time")
	ErrInvalidTimeFormat = errors.New("Invalid time format")
)

type ShiftTime struct {
	T time.Time
}

func NewShiftTime(t time.Time) ShiftTime {
	return ShiftTime{t}
}

func (d ShiftTime) String() string {
	h, m, _ := d.T.Clock()
	return fmt.Sprintf("%02d:%02d", h, m)
}

func (d *ShiftTime) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}

	t, err := time.Parse("15:04", s)
	if err != nil {
		return ErrInvalidTimeFormat
	}

	d.T = t
	return nil
}

func (d *ShiftTime) Validate() error {
	today := time.Now()

	before := d.T.Before(today)
	if before {
		return ErrInvalidTime
	}

	return nil
}
