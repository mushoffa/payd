package valueobject

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var (
	ErrInvalidDate       = errors.New("Invalid date")
	ErrInvalidDateFormat = errors.New("Invalid date format")
)

type ShiftDate struct {
	T time.Time
}

func NewShiftDate(t time.Time) *ShiftDate {
	return &ShiftDate{t}
}

func (d *ShiftDate) String() string {
	y, m, day := d.T.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, day)
}

func (d *ShiftDate) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return ErrInvalidDateFormat
	}

	d.T = t
	return nil
}

func (d *ShiftDate) Validate() error {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	before := d.T.Before(today)
	if before {
		return ErrInvalidDate
	}

	return nil
}
