package models

import (
	"gorm.io/gorm"
	"time"
)

type Income struct {
	gorm.Model
	UserID    uint      `gorm:"not null"`
	Source    string    `gorm:"not null"`
	Amount    float64   `gorm:"not null"`
	Date      time.Time `gorm:"not null"`
	Recurring bool      `gorm:"default:false"`
	Note      string
}
