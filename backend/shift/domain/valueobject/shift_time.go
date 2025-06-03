package valueobject

import (
	"encoding/json"
	"fmt"
	"time"

	"payd/domain/error"
	vo "payd/domain/valueobject"
)

type ShiftTime vo.Time

func NewShiftTime(t string) (ShiftTime, domain.Error) {
	_time, err := parseTime(t)
	if err != nil {
		return ShiftTime(time.Time{}), err
	}

	return ShiftTime(_time), nil
}

func (e ShiftTime) String() string {
	hour, minute, _ := e.Time().Clock()
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

func (e ShiftTime) Time() time.Time {
	return time.Time(e)
}

func (e ShiftTime) MarshalJSON() ([]byte, error) {
	jsonString := fmt.Sprintf("\"%s\"", e.String())
	return []byte(jsonString), nil
}

func (e *ShiftTime) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}

	time, err := parseTime(s)
	if err != nil {
		return err
	}

	*e = ShiftTime(time)

	return nil
}

func parseTime(t string) (vo.Time, domain.Error) {
	_time, err := vo.NewTime(t)
	if err != nil {
		return vo.Time(time.Time{}), domain.BadRequest(err.Error())
	}

	return _time, nil
}
