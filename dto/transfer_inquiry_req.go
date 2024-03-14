package dto

type TransferInquiryRequest struct {
	DofNumber string  `json:"dof_number" validate:"required"`
	Amount    float64 `json:"amount" validate:"required"`
}
