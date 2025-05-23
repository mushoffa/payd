package regex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type date_test struct {
	in  string
	out error
}

func Test_Date_ddmmyyyy(t *testing.T) {
	// Given
	tests := []date_test{
		{
			"01/12/2025",
			nil,
		},
		{
			"01-12-2025",
			nil,
		},
		{
			"01/12-2025",
			ErrInvalidDateFormat,
		},
		{
			"01-12/2025",
			ErrInvalidDateFormat,
		},
	}

	for i := range tests {
		tc := tests[i]
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			t.Parallel()

			// When
			err := Date(tc.in)

			// Then
			assert.Equal(t, tc.out, err)
		})
	}
}

func Test_Date_yyyymmdd(t *testing.T) {
	// Given
	tests := []date_test{
		{
			"2025-12-12",
			nil,
		},
		{
			"2025/12/21",
			nil,
		},
		{
			"2025-12/31",
			ErrInvalidDateFormat,
		},
		{
			"2025/12-31",
			ErrInvalidDateFormat,
		},
	}

	for i := range tests {
		tc := tests[i]
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			t.Parallel()

			// When
			err := Date(tc.in)

			// Then
			assert.Equal(t, tc.out, err)
		})
	}
}
