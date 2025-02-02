package models

import (
	"gorm.io/gorm"
	"time"
)

type Expense struct {
	gorm.Model
	UserID     uint      `gorm:"not null"`
	CategoryID uint      `binding:"required" gorm:"not null"`
	Amount     float64   `binding:"required" gorm:"not null"`
	Date       time.Time `binding:"required" gorm:"not null"`
	Note       string    `binding:"required"`
	Recurring  bool      `gorm:"default:false"`
}
