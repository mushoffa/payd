package regex

import (
	"errors"
	"fmt"
	"regexp"
)

// List of regex for date formats
const (
	// DD/MM/YYYY regex pattern.
	ddmmyyyy_s = `((0[1-9]|[12][0-9]|3[01])\/(0[1-9]|1[0,1,2])\/(19|20)\d{2})`

	// DD-MM-YYYY regex pattern.
	//ddmmyyyy_d = `(^(\d{4})-(0[1-9]|1[0-2]|[1-9])-([1-9]|0[1-9]|[1-2]\d|3[0-1])$)`
	ddmmyyyy_d = `((0[1-9]|[12][0-9]|3[01])\-(0[1-9]|1[0,1,2])\-(19|20)\d{2})`

	// YYYY/MM/DD regex pattern.
	yyyymmdd_s = `([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))`

	// YYYY-MM-DD regex pattern.
	yyyymmdd_d = `(^(\d{4})/(0[1-9]|1[0-2]|[1-9])/([1-9]|0[1-9]|[1-2]\d|3[0-1])$)`
)

var (
	ErrInvalidDateFormat = errors.New("Invalid date format")
)

func Date(date string) error {
	regex := fmt.Sprintf("%s|%s|%s|%s", ddmmyyyy_s, ddmmyyyy_d, yyyymmdd_s, yyyymmdd_d)
	pattern := regexp.MustCompile(regex)
	valid := pattern.MatchString(date)

	if !valid {
		return ErrInvalidDateFormat
	}
	return nil
}
