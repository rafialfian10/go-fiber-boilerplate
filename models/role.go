package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Role string `gorm:"unique"`
}
