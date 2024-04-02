package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/helper"
	"github.com/midoon/e-wallet-go-app-v1/util"
	"github.com/redis/go-redis/v9"
)

type transactionService struct {
	transactionRepository  domain.TransactionRepository
	accountRepository      domain.AccountRepository
	notificationRepository domain.NotificationRepository
	rdb                    *redis.Client
	validate               *validator.Validate
}

func NewTransactionService(transactionRepository domain.TransactionRepository, accountRepository domain.AccountRepository, notificationRepository domain.NotificationRepository, rdb *redis.Client, validate *validator.Validate) domain.TransactionService {
	return &transactionService{
		transactionRepository:  transactionRepository,
		accountRepository:      accountRepository,
		notificationRepository: notificationRepository,
		rdb:                    rdb,
		validate:               validate,
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
	err := t.validate.Struct(req)
	if err != nil {
		return helper.ErrValidation
	}

	myAccId := ctx.Value("x-user-id").(string)
	myAccount, err := t.accountRepository.FindByUserId(ctx, myAccId)
	if err != nil {
		return helper.ErrAccountNotFound
	}

	if myAccount == (domain.Account{}) {
		return helper.ErrAccountNotFound
	}

	if ok := util.CheckPasswordHash(req.UserPin, myAccount.Pin); !ok {
		return helper.ErrAccessDenied
	}

	tranferData, err := t.rdb.Get(ctx, req.InquiryKey).Result()
	if err != nil {
		return err
	}

	var inqReq dto.TransferInquiryRequest
	if err := json.Unmarshal([]byte(tranferData), &inqReq); err != nil {
		return helper.ErrInquiryNotFound
	}

	if inqReq == (dto.TransferInquiryRequest{}) {
		return helper.ErrInquiryNotFound
	}

	dofAccount, err := t.accountRepository.FindByAccNum(ctx, inqReq.DofNumber)
	if err != nil {
		return helper.ErrAccountNotFound
	}

	if dofAccount == (domain.Account{}) {
		return helper.ErrAccountNotFound
	}

	debit := domain.Transaction{
		AccountId:       myAccount.ID,
		SofNumber:       myAccount.AccountNumber,
		DofNumber:       dofAccount.AccountNumber,
		Amount:          inqReq.Amount,
		TransactionType: "D",
	}

	credit := domain.Transaction{
		AccountId:       dofAccount.ID,
		SofNumber:       myAccount.AccountNumber,
		DofNumber:       dofAccount.AccountNumber,
		Amount:          inqReq.Amount,
		TransactionType: "C",
	}

	err = t.transactionRepository.Insert(ctx, &debit, &credit)
	if err != nil {
		return err
	}

	err = t.rdb.Del(ctx, req.InquiryKey).Err()
	if err != nil {
		return err
	}

	// create notification data
	go t.notificationAfterTransfer(myAccount, dofAccount, inqReq.Amount)

	return nil
}

// tidak dibuatkan interface karena hanya dipakai di interenal (tidak digunakan sebagai API)
func (t *transactionService) notificationAfterTransfer(sender domain.Account, reciever domain.Account, amount float64) {
	senderNotificaton := domain.Notification{
		Title:     "Transfer Berhasil",
		Body:      fmt.Sprintf("Transfer senilai %.2f ke %s telah berhasil", amount, reciever.AccountNumber),
		Status:    1,
		IsRead:    0,
		AccountId: sender.ID,
	}

	recieverNotification := domain.Notification{
		Title:     "Dana Diterima",
		Body:      fmt.Sprintf("Dana diterima senilai %.2f dari %s", amount, sender.AccountNumber),
		Status:    1,
		IsRead:    0,
		AccountId: reciever.ID,
	}

	_ = t.notificationRepository.Insert(context.Background(), &senderNotificaton)
	_ = t.notificationRepository.Insert(context.Background(), &recieverNotification)
}
