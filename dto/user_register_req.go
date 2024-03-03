package dto

type UserRegisterRequest struct {
	Username      string `json:"username" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required,min=8"`
	AccountNumber string `json:"account_number" validate:"required,min=8"`
	Pin           string `json:"pin" validate:"required,min=4,max=4"`
}
