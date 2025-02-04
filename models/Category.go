package models

import (
	"time"
)

type Category struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `json:"name" binding:"required" gorm:"not null"`
	Icon      string     `json:"icon" binding:"required" gorm:"not null"`
	UserID    uint       `json:"userId" gorm:"not null"`
}
