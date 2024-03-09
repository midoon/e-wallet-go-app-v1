package dto

type InquirryKey struct {
	InquiryKey string `json:"inquiry_key"`
}

type TransferInquiryResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    InquirryKey `json:"data"`
}
