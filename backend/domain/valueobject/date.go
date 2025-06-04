package valueobject

import (
	"time"

	"payd/domain/regex"
)

type Date time.Time

type DateFormat int

// List of date formats
const (
	// DDMMYYYY_S represents the date format with '/' separator: DD/MM/YYYY.
	DDMMYYYY_S DateFormat = iota

	// DDMMYYYY_D represents the date format with '-' separator: DD-MM-YYYY.
	DDMMYYYY_D

	// YYYYMMDD_S represents the date format with '/' separator: YYYY/MM/DD.
	YYYYMMDD_S

	// YYYYMMDD_S represents the date format with '-' separator: YYYY-MM-DD.
	YYYYMMDD_D
)

var dateFormats = map[DateFormat]string{
	DDMMYYYY_S: "02/01/2006",
	DDMMYYYY_D: "02-01-2006",
	YYYYMMDD_S: "2006/01/02",
	YYYYMMDD_D: "2006-01-02",
}

func NewDate(date string, format DateFormat) (Date, error) {
	if err := regex.Date(date); err != nil {
		return Date(time.Time{}), err
	}

	_date := parse(date, format)

	return Date(_date), nil
}

func (e Date) Time() time.Time {
	return time.Time(e)
}

func (e DateFormat) String() string {
	return dateFormats[e]
}

func parse(date string, format DateFormat) time.Time {
	_date, _ := time.Parse(format.String(), date)

	return _date
}
