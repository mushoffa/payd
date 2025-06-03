package valueobject

import (
	"encoding/json"
	"fmt"
	"time"

	"payd/domain/error"
	vo "payd/domain/valueobject"
	shift "payd/shift/domain/error"
)

type ShiftDate vo.Date

func NewShiftDate(date string) (ShiftDate, domain.Error) {
	_date, err := parseDate(date)
	if err != nil {
		return ShiftDate(time.Time{}), err
	}

	if err := validateShiftDate(_date.Time()); err != nil {
		return ShiftDate(time.Time{}), err
	}

	return ShiftDate(_date), nil
}

func (e ShiftDate) String() string {
	y, m, d := e.Time().Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

func (e ShiftDate) Time() time.Time {
	return time.Time(e)
}

func (e ShiftDate) MarshalJSON() ([]byte, error) {
	jsonString := fmt.Sprintf("\"%s\"", e.String())
	return []byte(jsonString), nil
}

func (e *ShiftDate) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}

	date, err := parseDate(s)
	if err != nil {
		return err
	}

	*e = ShiftDate(date)

	return nil
}

func parseDate(date string) (vo.Date, domain.Error) {
	_date, err := vo.NewDate(date, vo.YYYYMMDD_D)
	if err != nil {
		return vo.Date(time.Time{}), domain.BadRequest(err.Error())
	}

	return _date, nil
}

func validateShiftDate(date time.Time) domain.Error {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	before := date.Before(today)
	if before {
		return shift.ErrInvalidShiftDate
	}

	return nil
}
