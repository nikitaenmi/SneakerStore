package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Namebrand string
	Size      int
	UserID    string
}
