package domain

import "time"

type Store struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	CreateAt time.Time
	UpdateAt time.Time
}
