package dto

type TransferInquiryRequest struct {
	AccountNumber string  `json:"account_number" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
}
