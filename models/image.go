package models

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID         uint   `gorm:"primaryKey"`
	Title      string `gorm:"type:varchar(255);not null"`
	DisasterID uint
	Disaster   Disaster `gorm:"foreignKey:DisasterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
