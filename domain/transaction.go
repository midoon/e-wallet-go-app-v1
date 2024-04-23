package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"gorm.io/gorm"
)

type Transaction struct {
	ID              string    `gorm:"primary_key;column:id"`
	SofNumber       string    `gorm:"column:sof_number"`
	DofNumber       string    `gorm:"column:dof_number"`
	Amount          float64   `gorm:"column:amount"`
	TransactionType string    `gorm:"column:transaction_type"`
	AccountId       string    `gorm:"column:account_id"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (t *Transaction) BeforeCreate(db *gorm.DB) error {
	if t.ID == "" {
		id := uuid.New().String()
		t.ID = id
	}

	return nil
}

type TransactionRepository interface {
	Insert(ctx context.Context, debit *Transaction, credit *Transaction) error
	InsertFromMidtrans(ctx context.Context, transaction *Transaction) error
}

type TransactionService interface {
	TransferInquiry(ctx context.Context, req dto.TransferInquiryRequest) (dto.InquirryKey, error)
	TranferExecute(ctx context.Context, req dto.TransferExecuteRequest) error
}
