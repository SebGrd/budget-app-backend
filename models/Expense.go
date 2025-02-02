package models

import (
	"gorm.io/gorm"
	"time"
)

type Expense struct {
	gorm.Model
	UserID     uint      `gorm:"not null"`
	CategoryID uint      `gorm:"not null"`
	Amount     float64   `gorm:"not null"`
	Date       time.Time `gorm:"not null"`
	Note       string
	Recurring  bool `gorm:"default:false"`
}
