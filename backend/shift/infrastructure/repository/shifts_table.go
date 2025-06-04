package repository

import (
	"time"
)

type ShiftTable struct {
	ID           int       `db:"id"`
	Created      time.Time `db:"created"`
	Updated      time.Time `db:"updated"`
	Date         time.Time `db:"date"`
	StartTime    time.Time `db:"start_time"`
	EndTime      time.Time `db:"end_time"`
	Role         string    `db:"role"`
	Location     string    `db:"location"`
	Assigned     bool      `db:"assigned"`
	EmployeeName *string   `db:"employee_name"`
	EmployeeID   *int      `db:"employee_id"`
}
