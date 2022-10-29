package models

import "github.com/uadmin/uadmin"

type Employee struct {
	uadmin.Model
	FirstName    string
	LastName     string
	Designation  string
	EmployeeType EmployeeType
}

type EmployeeType int

func (EmployeeType) RankAndFile() int {
	return 1
}
func (EmployeeType) Supervisory() int {
	return 2
}
func (EmployeeType) Managerial() int {
	return 3
}
func (EmployeeType) Executive() int {
	return 4
}
