package repository

import (
	"context"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"gorm.io/gorm"
)

type transactionRespository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &transactionRespository{
		db: db,
	}
}

func (t *transactionRespository) Insert(ctx context.Context, transaction *domain.Transaction) error {
	panic("not implemented") // TODO: Implement
}
