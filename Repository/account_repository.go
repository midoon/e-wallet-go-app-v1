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

func (a *accountRepository) Update(ctx context.Context, account *domain.Account, accountId string) error {
	panic("not implemented") // TODO: Implement
}

func (a *accountRepository) FindByAccNum(ctx context.Context, accNum string) (domain.Account, error) {
	panic("not implemented") // TODO: Implement
}

func (a *accountRepository) FindById(ctx context.Context, accountId string) (domain.Account, error) {
	panic("not implemented") // TODO: Implement
}
