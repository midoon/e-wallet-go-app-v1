package dto

type TopUpRequest struct {
	Amount float64 `json:"amount"`
	UserId string  `json:"-"`
}
