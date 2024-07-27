package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID        string
	Namebrand string
}
