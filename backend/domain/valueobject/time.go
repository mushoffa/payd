package valueobject

import (
	"time"

	"payd/domain/regex"
)

type Time time.Time

func NewTime(t string) (Time, error) {
	if err := regex.Time(t); err != nil {
		return Time(time.Time{}), err
	}

	_time := parseTime(t)

	return Time(_time), nil
}

func parseTime(t string) time.Time {
	_time, _ := time.Parse("15:04", t)

	return _time
}
