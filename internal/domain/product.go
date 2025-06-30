package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"image_url"`
	StoreID     uint    `json:"store_id"`
	CategoryID  uint    `json:"category_id"`

	Store    Store    `json:"store" gorm:"foreignKey:StoreID;reference:ID"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID;reference:ID"`
}
