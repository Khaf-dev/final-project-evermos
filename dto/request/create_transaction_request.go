package request

import "github.com/go-playground/validator/v10"

type TransactionItem struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,min=1"`
}

type CreateTransactionRequest struct {
	StoreID uint              `json:"store_id" validate:"required"`
	Items   []TransactionItem `json:"items" validate:"required,dive"`
}

func (r *CreateTransactionRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
