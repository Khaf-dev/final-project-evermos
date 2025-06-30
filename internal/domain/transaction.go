package domain

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID  uint                `json:"user_id"`
	StoreID uint                `json:"store_id"`
	Total   float64             `json:"total"`
	Status  string              `gorm:"default:'pending'"`
	Details []TransactionDetail `json:"details" gorm:"foreignKey:TransactionID"`

	User  User  `json:"user" gorm:"foreignKey:UserID"`
	Store Store `json:"store" gorm:"foreignKey:StoreID"`
}

type TransactionDetail struct {
	gorm.Model
	TransactionID uint        `json:"transaction_id"`
	Transaction   Transaction `json:"-"`
	ProductID     uint        `json:"product_id"`
	Product       Product     `json:"product" gorm:"foreignKey:ProductID"`
	Quantity      int         `json:"quantity"`
	Price         float64     `json:"price"`
	Subtotal      float64     `json:"sub_total"`
}
