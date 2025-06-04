package valueobject

import (
	"fmt"
	"testing"
	"time"

	"payd/domain/regex"

	"github.com/stretchr/testify/assert"
)

type date_test struct {
	in     string
	format DateFormat
	out    error
}

func Test_Date_ddmmyyyy(t *testing.T) {
	// Given
	tests := []date_test{
		{
			"01/12/2025",
			DDMMYYYY_S,
			nil,
		},
		{
			"01-12-2025",
			DDMMYYYY_D,
			nil,
		},
		{
			"01/12-2025",
			DDMMYYYY_S,
			regex.ErrInvalidDateFormat,
		},
		{
			"01-12/2025",
			DDMMYYYY_D,
			regex.ErrInvalidDateFormat,
		},
	}

	for i := range tests {
		tc := tests[i]
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			t.Parallel()

			// When
			date, err := NewDate(tc.in, tc.format)
			t.Log(time.Time(date).GoString())

			// Then
			assert.Equal(t, tc.out, err)
		})
	}
}
