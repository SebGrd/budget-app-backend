package models

import (
	"time"
)

type Income struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	UserID    uint       `json:"userId" gorm:"not null"`
	Source    string     `json:"source" gorm:"not null"`
	Amount    float64    `json:"amount" gorm:"not null"`
	Date      time.Time  `json:"date" gorm:"not null"`
	Recurring bool       `json:"recurring" gorm:"default:false"`
	Note      string     `json:"note"`
}
