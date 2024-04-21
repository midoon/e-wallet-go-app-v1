package repository

import (
	"context"
	"log"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"gorm.io/gorm"
)

type topupRepository struct {
	db *gorm.DB
}

func NewTopUpRepository(db *gorm.DB) domain.TopupRepository {
	return &topupRepository{
		db: db,
	}
}

// FindById implements domain.TopupRepository.
func (t *topupRepository) FindById(ctx context.Context, id string) (domain.Topup, error) {
	topup := domain.Topup{}
	err := t.db.WithContext(ctx).Where("id = ?", id).Take(&topup).Error
	if err != nil {
		log.Println(err)
		return domain.Topup{}, err
	}
	return topup, nil
}

// Insert implements domain.TopupRepository.
func (t *topupRepository) Insert(ctx context.Context, topup *domain.Topup) error {
	err := t.db.WithContext(ctx).Create(topup).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Update implements domain.TopupRepository.
func (t *topupRepository) Update(ctx context.Context, topup *domain.Topup) error {
	err := t.db.WithContext(ctx).Where("id = ?", topup.ID).Updates(domain.Topup{
		Status:  topup.Status,
		Amount:  topup.Amount,
		SnapUrl: topup.SnapUrl,
	}).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
