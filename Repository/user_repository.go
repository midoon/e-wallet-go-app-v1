package repository

import (
	"context"
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

func (u *userRepository) FindById(ctx context.Context, id string) (domain.User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user := domain.User{}
	d := u.db.WithContext(ctx).Where("email = ?", email).Take(&user)
	return user, d.Error
}

func (u *userRepository) Insert(ctx context.Context, user *domain.User) error {
	d := u.db.WithContext(ctx).Create(user)
	if d.Error != nil {
		log.Fatal(d.Error)
	}
	return d.Error
}

func (u *userRepository) Update(ctx context.Context, user *domain.User) error {
	panic("not implemented") // TODO: Implement
}
