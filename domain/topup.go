package domain

import (
	"context"
	"time"

	"github.com/midoon/e-wallet-go-app-v1/dto"
)

type Topup struct {
	ID        string    `gorm:"column:id;primary_key"`
	Status    int8      `gorm:"column:status"`
	Amount    float64   `gorm:"column:amount"`
	SnapUrl   string    `gorm:"column:snap_url"`
	UserId    string    `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

// ID akakn dibuat maunial

type TopupRepository interface {
	FindById(ctx context.Context, id string) (Topup, error)
	Insert(ctx context.Context, topup *Topup) error
	Update(ctx context.Context, topup *Topup) error
}

type TopupService interface {
	ConfirmedTopUp(ctx context.Context, id string) error
	InitializeTopUp(ctx context.Context, req dto.TopUpRequest) (dto.TopUpResponse, error)
}
