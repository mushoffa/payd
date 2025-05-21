package valueobject

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type shift_date struct {
	Date ShiftDate `json:"date"`
}

func Test_ShiftDate_Validate_Today(t *testing.T) {
	// Given
	today := time.Now()
	shift := NewShiftDate(today)

	// When
	err := shift.Validate()

	// Then
	assert.Nil(t, err)
}

func Test_ShiftDate_Validate_Yesterday(t *testing.T) {
	// Given
	yesterday := time.Now().Add(-1 * time.Hour *24)
	shift := NewShiftDate(yesterday)

	// When
	err := shift.Validate()

	// Then
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrInvalidDate)
}

func Test_ShiftDate_UnmarshalJSON_Success(t *testing.T) {
	// Given
	var date shift_date
	todayDate := getDate()
	jsonPayload := []byte(fmt.Sprintf(`{"date":"%s"}`, todayDate))

	// When
	err := json.Unmarshal(jsonPayload, &date)

	// Then
	assert.Nil(t, err)
	assert.Equal(t, todayDate, date.Date.String())
}

func getDate() string {
	y, m, d := time.Now().Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}