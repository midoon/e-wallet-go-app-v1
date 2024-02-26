package repository

import (
	"context"
	"log"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"gorm.io/gorm"
)

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) domain.TokenRepository {
	return &tokenRepository{
		db: db,
	}
}

func (t *tokenRepository) Insert(ctx context.Context, token *domain.Token) error {
	err := t.db.WithContext(ctx).Create(token).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (t *tokenRepository) FindByUserId(ctx context.Context, userId string) (domain.Token, error) {
	token := domain.Token{}
	err := t.db.WithContext(ctx).Where("user_id = ?", userId).Take(&token).Error
	if err != nil {
		log.Println(err)
		return domain.Token{}, err
	}
	return token, nil
}

func (t *tokenRepository) CountByUserId(ctx context.Context, userId string) (int64, error) {
	var count int64
	err := t.db.WithContext(ctx).Model(&domain.Token{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return count, nil
}

func (t *tokenRepository) Delete(ctx context.Context, userId string) error {
	err := t.db.WithContext(ctx).Delete(&domain.Token{}, "user_id = ?", userId).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
