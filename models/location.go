package models

import "github.com/uadmin/uadmin"

type Location struct {
	uadmin.Model
	Name        string
	Description string
}
