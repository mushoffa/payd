package valueobject

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShiftDate_validate(t *testing.T) {

	// Given
	tests := []struct {
		description string
		in          string
		out         error
	}{
		{
			in:  "2025-01-21",
			out: ErrInvalidShiftDate,
		},
	}

	for i := range tests {
		tc := tests[i]
		t.Run(fmt.Sprintf("%v", tc.description), func(t *testing.T) {
			t.Parallel()

			// When
			date, err := NewShiftDate(tc.in)
			t.Log(date)

			// Then
			assert.Equal(t, tc.out, err)
		})
	}
}
