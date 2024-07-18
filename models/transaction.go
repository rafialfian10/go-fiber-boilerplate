package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID              uint `gorm:"primaryKey"`
	UserID          uuid.UUID
	User            User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DisasterID      uint
	Disaster        Disaster  `gorm:"foreignKey:DisasterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status          string    `json:"status"  gorm:"type:varchar(25)"`
	TransactionDate time.Time `gorm:"type:timestamp;not null"`
	Token           string    `json:"token" gorm:"type: varchar(255)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
