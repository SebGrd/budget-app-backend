package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string `binding:"required" gorm:"unique;not null"`
	Icon   string `binding:"required" gorm:"not null"`
	UserID uint   `gorm:"not null"`
}
