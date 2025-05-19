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
