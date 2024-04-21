package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
)

type topupService struct {
	notificationRepository domain.NotificationRepository
	topupRepository        domain.TopupRepository
	midtransService        domain.MidtransService
	accountRepository      domain.AccountRepository
}

func NewTopupService(notificationRepository domain.NotificationRepository, topupRepository domain.TopupRepository, midtransService domain.MidtransService, accountRepository domain.AccountRepository) domain.TopupService {
	return &topupService{
		notificationRepository: notificationRepository,
		topupRepository:        topupRepository,
		midtransService:        midtransService,
		accountRepository:      accountRepository,
	}
}

// ConfirmedTopUp implements domain.TopupService.
func (t *topupService) ConfirmedTopUp(ctx context.Context, id string) error {
	topUP, err := t.topupRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	if topUP == (domain.Topup{}) {
		return errors.New("top-up not found")
	}

	account, err := t.accountRepository.FindByUserId(ctx, topUP.UserId)
	if err != nil {
		return err
	}

	if account == (domain.Account{}) {
		return errors.New("account not found")
	}

	account.Balance += topUP.Amount
	err = t.accountRepository.Update(ctx, &account, account.ID)
	if err != nil {
		return err
	}

	notif := domain.Notification{
		Title:     "top-up",
		Body:      fmt.Sprintf("sukse topup saldo sebesar %.2f", topUP.Amount),
		Status:    1,
		IsRead:    0,
		AccountId: account.ID,
	}
	_ = t.notificationRepository.Insert(ctx, &notif)
	return err
}

func (t *topupService) InitializeTopUp(ctx context.Context, req dto.TopUpRequest) (dto.TopUpResponse, error) {
	topUp := domain.Topup{
		ID:     uuid.NewString(),
		UserId: req.UserId,
		Status: 0,
		Amount: req.Amount,
	}

	err := t.midtransService.GenerateSnapUrl(ctx, &topUp)
	if err != nil {
		return dto.TopUpResponse{}, err
	}

	return dto.TopUpResponse{
		SnapUrl: topUp.SnapUrl,
	}, nil
}
