package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/helper"
	"github.com/midoon/e-wallet-go-app-v1/util"
	"github.com/redis/go-redis/v9"
)

type transactionService struct {
	transactionRepository domain.TransactionRepository
	accountRepository     domain.AccountRepository
	rdb                   *redis.Client
	validate              *validator.Validate
}

func NewTransactionService(transactionRepository domain.TransactionRepository, accountRepository domain.AccountRepository, rdb *redis.Client, validate *validator.Validate) domain.TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
		rdb:                   rdb,
		validate:              validate,
	}
}

func (t *transactionService) TransferInquiry(ctx context.Context, req dto.TransferInquiryRequest) (dto.InquirryKey, error) {
	err := t.validate.Struct(req)
	if err != nil {
		return dto.InquirryKey{}, helper.ErrValidation
	}

	myUserId := ctx.Value("x-user-id").(string)
	myAccount, err := t.accountRepository.FindByUserId(ctx, myUserId)
	if err != nil {
		return dto.InquirryKey{}, err
	}
	if myAccount == (domain.Account{}) {
		return dto.InquirryKey{}, helper.ErrAccountNotFound
	}

	DofAccount, err := t.accountRepository.FindByAccNum(ctx, req.DofNumber)
	if err != nil {
		return dto.InquirryKey{}, err
	}
	if DofAccount == (domain.Account{}) {
		return dto.InquirryKey{}, helper.ErrAccountNotFound
	}

	if myAccount.Balance < req.Amount {
		return dto.InquirryKey{}, helper.ErrInsufficient
	}

	// generate random string && store data [randKey : JSON(req)] to redis
	inquiryKey, err := util.GenerateRandomString(30)
	if err != nil {
		return dto.InquirryKey{}, err
	}

	transferData, _ := json.Marshal(req)

	err = t.rdb.Set(ctx, inquiryKey, transferData, 1*time.Hour).Err()
	if err != nil {
		return dto.InquirryKey{}, err
	}

	return dto.InquirryKey{
		InquiryKey: inquiryKey,
	}, nil
}

func (t *transactionService) TranferExecute(ctx context.Context, req dto.TransferExecuteRequest) error {
	panic("not implemented") // TODO: Implement
}
