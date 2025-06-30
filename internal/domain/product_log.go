package domain

import "gorm.io/gorm"

type ProductLog struct {
	gorm.Model
	ProductID      uint
	ProductName    string
	PreviousStock  int
	PurchasedQty   int
	RemainingStock int
	TransactionID  uint
	Activity       string
	Detail         string
}
