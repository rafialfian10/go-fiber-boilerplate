package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID `gorm:"type:varchar(255);primaryKey"`
	FullName        string    `gorm:"type:varchar(255)"`
	Email           string    `gorm:"type:varchar(255);unique"`
	IsEmailVerified bool
	Password        string `gorm:"type:varchar(255)"`
	Phone           string `gorm:"type:varchar(255);unique"`
	IsPhoneVerified bool
	Gender          string `gorm:"type:text"`
	Address         string `gorm:"type:text"`
	RoleID          uint
	Role            Role
	Image           string `gorm:"type:varchar(255)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
