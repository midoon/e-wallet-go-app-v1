package dto

type TransferExecuteRequest struct {
	InquirryKey string `json:"inquiry_key" validate:"required"`
	UserPin     string `json:"user_pin" validate:"required"`
}
