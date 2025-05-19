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
