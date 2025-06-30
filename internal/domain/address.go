package domain

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID     uint   `json:"user_id"`
	Receiver   string `json:"receiver"`
	Phone      string `json:"phone"`
	Province   string `json:"province"`
	Regency    string `json:"regency"`
	District   string `json:"district"`
	Village    string `json:"village"`
	PostalCode string `json:"postal_code"`
	Detail     string `json:"detail"`
	IsPrimary  bool   `json:"is_primary"`
}
