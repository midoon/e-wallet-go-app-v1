package repository

import (
	"context"
	"log"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type transactionRespository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &transactionRespository{
		db: db,
	}
}

func (t *transactionRespository) Insert(ctx context.Context, debit *domain.Transaction, credit *domain.Transaction) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		var myAccount domain.Account
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&myAccount, "id = ?", debit.AccountId).Error
		if err != nil {
			return err
		}

		var dofAccount domain.Account
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&dofAccount, "id = ?", credit.AccountId).Error
		if err != nil {
			return err
		}

		// kalau parameter berupa pointer => *domain.Transaction, maka tidak perlu dijadikan pointer lagi ketika nilai parameter digunakan
		err = tx.Create(debit).Error
		if err != nil {
			return err
		}
		err = tx.Create(credit).Error
		if err != nil {
			return err
		}

		myAccount.Balance = myAccount.Balance - debit.Amount
		dofAccount.Balance = dofAccount.Balance + credit.Amount
		err = tx.Save(&myAccount).Error
		if err != nil {
			return err
		}
		err = tx.Save(&dofAccount).Error
		if err != nil {
			return err
		}

		return err
	})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (t *transactionRespository) InsertFromMidtrans(ctx context.Context, transaction *domain.Transaction) error {
	err := t.db.WithContext(ctx).Create(transaction).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
