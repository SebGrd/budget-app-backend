package models

import (
	"time"
)

type Expense struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
	UserID     uint       `json:"userId" gorm:"not null"`
	CategoryID uint       `json:"categoryId" binding:"required" gorm:"not null"`
	Amount     float64    `json:"amount" binding:"required" gorm:"not null"`
	Date       time.Time  `json:"date" binding:"required" gorm:"not null"`
	Note       string     `json:"note" binding:"required"`
	Recurring  bool       `json:"recurring" gorm:"default:false"`
}
