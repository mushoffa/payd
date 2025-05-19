package valueobject

import (
	"encoding/json"
	"fmt"
	"time"
)

type ShiftDate struct {
	T time.Time
}

func (d *ShiftDate) String() string {
	y, m, day := d.T.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, day)
}

func (d *ShiftDate) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to a string: %w", err)
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}

	d.T = t
	return nil
}
