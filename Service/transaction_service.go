package service

import (
	"context"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/redis/go-redis/v9"
)

type transactionService struct {
	transactionRepository domain.TransactionRepository
	rdb                   *redis.Client
}

func NewTransactionService(transactionRepository domain.TransactionRepository, rdb *redis.Client) domain.TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
		rdb:                   rdb,
	}
}

func (t *transactionService) TransferInquiry(ctx context.Context, req dto.TransferInquiryRequest) (dto.InquirryKey, error) {
	panic("not implemented") // TODO: Implement
}

func (t *transactionService) TranferExecute(ctx context.Context, req dto.TransferExecuteRequest) error {
	panic("not implemented") // TODO: Implement
}
