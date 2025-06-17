package domain

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Phone    string `gorm:"uniqueIndex; not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:user"`
	CreateAt time.Time
	UpdateAt time.Time
	Store    Store
}
