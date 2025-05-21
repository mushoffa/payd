package valueobject

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type shift_time struct {
	Time ShiftTime `json:"time"`
}

func Test_ShiftTime_Validate_Past_Hour(t *testing.T) {
	// Given
	today := time.Now().Add(-1 * time.Hour * 1)
	shift_time := NewShiftTime(today)

	// When
	err := shift_time.Validate()

	// Then
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrInvalidTime)
}

func Test_ShiftTime_Validate_Past_Minute(t *testing.T) {
	// Given
	today := time.Now().Add(-1 * time.Minute * 1)
	shift_time := NewShiftTime(today)

	// When
	err := shift_time.Validate()

	// Then
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrInvalidTime)
}

func Test_ShiftTime_UnmarshalJSON_Success(t *testing.T) {
	// Given
	var time shift_time
	todayTime := getTime()
	jsonPayload := []byte(fmt.Sprintf(`{"time":"%s"}`, todayTime))

	// When
	err := json.Unmarshal(jsonPayload, &time)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, todayTime, time.Time.String())
}

func getTime() string {
	h, m, _ := time.Now().Clock()
	return fmt.Sprintf("%02d:%02d", h, m)
}
