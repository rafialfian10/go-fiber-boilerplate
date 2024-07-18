package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Disaster struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uuid.UUID
	User         User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CategoryID   uint
	Category     Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title        string    `gorm:"type:varchar(255);not null"`
	Description  string    `gorm:"type:text;not null"`
	Location     string    `gorm:"type:varchar(255);not null"`
	Date         time.Time `gorm:"type:timestamp;not null"`
	Donate       int       `gorm:"not null"`
	DonateTarget int       `gorm:"not null"`
	Image        string    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
