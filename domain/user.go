package domain

import (
	"context"
	"time"

	"github.com/midoon/e-wallet-go-app-v1/dto"
)

type User struct {
	ID        string    `gorm:"primary_key;column:id"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email;uniqueIndex"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime"`
	Token     Token     `gorm:"foreignKey:user_id;referenceid"`
}

type UserRepository interface {
	FindById(ctx context.Context, userId string) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	CountByEmail(ctx context.Context, email string) (int64, error)
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User, userId string) error
}

type UserService interface {
	Register(ctx context.Context, req dto.UserRegisterRequest) error
}
