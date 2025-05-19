package valueobject

import (
	"encoding/json"
	"fmt"
	"time"
)

type ShiftTime struct {
	T time.Time
}

func (d *ShiftTime) String() string {
	h, m, _ := d.T.Clock()
	return fmt.Sprintf("%02d:%02d", h, m)
}

func (d *ShiftTime) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}

	t, err := time.Parse("15:04", s)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}

	d.T = t
	return nil
}
