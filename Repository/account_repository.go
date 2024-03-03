package repository

import (
	"context"
	"log"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) domain.AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (a *accountRepository) Insert(ctx context.Context, account *domain.Account) error {
	err := a.db.WithContext(ctx).Create(account).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
