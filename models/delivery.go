package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Deliveries struct {
	uadmin.Model
	Code       string
	Employee   Employee
	EmployeeID uint
	Status     Status
	StartTime  time.Time
	EndTime    time.Time
	Duration   time.Duration
}

type Status int

func (Status) Open() int {
	return 1
}
func (Status) Ongoing() int {
	return 2
}
func (Status) Confirmed() int {
	return 3
}
func (Status) Unconfirmed() int {
	return 4
}
