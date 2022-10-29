package models

import "github.com/uadmin/uadmin"

type Tags struct {
	uadmin.Model
	Code       string
	Name       string
	Location   Location
	LocationID uint
}

func (t *Tags) String() string {
	return t.Code + " " + t.Name
}
