package domain

import (
	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	Name   string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}
