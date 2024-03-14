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
	err := a.db.WithContext(ctx).Where("id = ?", accountId).Updates(domain.Account{
		Balance: account.Balance,
	}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (a *accountRepository) FindByAccNum(ctx context.Context, accNum string) (domain.Account, error) {
	account := domain.Account{}
	err := a.db.WithContext(ctx).Where("account_number = ?", accNum).Take(&account).Error
	if err != nil {
		log.Println(err)
		return domain.Account{}, err
	}
	return account, nil

}

func (a *accountRepository) FindById(ctx context.Context, accountId string) (domain.Account, error) {
	account := domain.Account{}
	err := a.db.WithContext(ctx).Where("id = ?", accountId).Take(&account).Error
	if err != nil {
		log.Println(err)
		return domain.Account{}, err
	}
	return account, nil
}

func (a *accountRepository) FindByUserId(ctx context.Context, userID string) (domain.Account, error) {
	account := domain.Account{}
	err := a.db.WithContext(ctx).Where("user_id = ?", userID).Take(&account).Error
	if err != nil {
		log.Println(err)
		return domain.Account{}, err
	}
	return account, nil
}
