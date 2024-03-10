package repository

import (
	"context"
	"log"

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
	err := t.db.WithContext(ctx).Create(transaction).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
