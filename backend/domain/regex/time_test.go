package regex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type time_test struct {
	in  string
	out error
}

func Test_Time_ValidFormat(t *testing.T) {
	var tests [60]time_test

	hour := 0
	for i := 0; i <= 59; i++ {
		_time := fmt.Sprintf("%02d:%02d", hour, i)
		if i < 23 {
			hour++
		}

		tests[i] = time_test{_time, nil}
	}

	for j := range tests {
		tc := tests[j]
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			t.Parallel()

			// When
			err := Time(tc.in)

			// Then
			assert.Equal(t, tc.out, err)
		})
	}
}
