package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Lastname string
	Size     int
}
