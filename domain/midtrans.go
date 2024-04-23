package domain

import "context"

type MidtransService interface {
	GenerateSnapUrl(ctx context.Context, t *Topup) error
	VerifyPayment(ctx context.Context, orderId string) (bool, error)
}
