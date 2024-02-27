package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) FindById(ctx context.Context, userId string) (domain.User, error) {
	user := domain.User{}
	err := u.db.WithContext(ctx).Where("id = ?", userId).Take(&user).Error
	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}
	return user, nil
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user := domain.User{}
	err := u.db.WithContext(ctx).Where("email = ?", email).Take(&user).Error
	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}
	return user, nil
}

func (u *userRepository) CountByEmail(ctx context.Context, email string) (int64, error) {
	var count int64
	err := u.db.WithContext(ctx).Model(&domain.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func (u *userRepository) Insert(ctx context.Context, user *domain.User) error {
	err := u.db.WithContext(ctx).Create(user).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (u *userRepository) Update(ctx context.Context, user *domain.User, userId string) error {
	// not using db.Model because the struc and Model used is same
	err := u.db.WithContext(ctx).Where("id = ?", userId).Updates(domain.User{
		Username: user.Username,
		Email:    user.Email,
	}).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
