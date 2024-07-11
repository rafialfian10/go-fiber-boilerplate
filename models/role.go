package models

import (
	"gorm.io/gorm"
)

// Role
type Role struct {
	gorm.Model
	Role string `gorm:"unique"`
}
