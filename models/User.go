package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Username  string     `json:"username" binding:"required" gorm:"unique;not null"`
	Password  string     `json:"password" binding:"required" gorm:"not null"`
}
