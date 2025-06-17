package request

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validated:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}
