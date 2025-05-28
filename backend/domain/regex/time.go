package regex

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	// HHmm represents a 24-hour time format with optional leading 0
	HHmm = `/^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$/`

	// HHmm_0 represents a 24-hour time format leading 0
	// ^(([0-1][0-9]|2[0-3]):[0-5][0-9](:[0-5][0-9])?)$
	HHmm_0 = `^(([0-1][0-9]|2[0-3]):[0-5][0-9])$`
)

var (
	ErrInvalidTimeFormat = errors.New("Invalid time format")
)

func Time(time string) error {
	regex := fmt.Sprintf("%s", HHmm_0)
	pattern := regexp.MustCompile(regex)
	valid := pattern.MatchString(time)

	if !valid {
		return ErrInvalidTimeFormat
	}
	return nil
}
