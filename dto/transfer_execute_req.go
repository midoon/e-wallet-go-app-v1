package dto

type TransferExecuteRequest struct {
	InquirryKey string `json:"inquiry_key" validate:"required"`
}
