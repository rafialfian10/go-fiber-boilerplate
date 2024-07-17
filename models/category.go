package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Category  string `json:"category" gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
