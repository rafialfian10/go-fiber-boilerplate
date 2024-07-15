package models

import (
	"time"
)

type Log struct {
	ID uint
	// UserID      uuid.UUID `gorm:"type:varchar(36);not null"`
	// User        Users     `gorm:"foreignKey:UserID;references:ID"`
	Date        time.Time
	IPAddress   string `gorm:"type:varchar(255)"`
	Host        string `gorm:"type:varchar(255)"`
	Path        string `gorm:"type:varchar(255)"`
	Method      string `gorm:"type:varchar(255)"`
	Body        string
	File        string
	ResposeTime float64
	StatusCode  int
}
