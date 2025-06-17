package request

type LoginRequest struct {
	EmailOrPhone string `json:"email_or_phone" validate:"required"`
	Password     string `json:"password" validate:"required"`
}
