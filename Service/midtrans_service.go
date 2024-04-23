package service

import (
	"context"
	"fmt"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/internal/config"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type midtransService struct {
	midtransConfig config.Midtrans
	envi           midtrans.EnvironmentType
}

func NewMidtransService(cnf *config.Config) domain.MidtransService {
	envi := midtrans.Sandbox
	if cnf.Midtrans.IsProd {
		envi = midtrans.Production
	}

	return &midtransService{
		midtransConfig: cnf.Midtrans,
		envi:           envi,
	}
}

// GenerateSnapUrl implements domain.MidtransService.
func (m *midtransService) GenerateSnapUrl(ctx context.Context, t *domain.Topup) error {
	// 2. Initiate Snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  t.ID,
			GrossAmt: int64(t.Amount),
		},
	}

	var client snap.Client
	client.New(m.midtransConfig.Key, m.envi)

	// 3. Request create Snap transaction to Midtrans
	snapResp, err := client.CreateTransaction(req)
	if err != nil {
		return err
	}

	t.SnapUrl = snapResp.RedirectURL
	return nil
}

// VerifyPayment implements domain.MidtransService.
func (m *midtransService) VerifyPayment(ctx context.Context, orderId string) (bool, error) {
	var client coreapi.Client

	client.New(m.midtransConfig.Key, m.envi)

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := client.CheckTransaction(orderId)
	if e != nil {
		return false, e
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
					return true, nil
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				// TODO set transaction status on your databaase to 'success'
				return true, nil
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
				fmt.Println("transaction pending")
			}
		}
	}
	return false, nil
}
