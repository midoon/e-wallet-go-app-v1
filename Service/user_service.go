package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/helper"
	"github.com/midoon/e-wallet-go-app-v1/util"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) Register(ctx context.Context, req dto.UserRegisterRequest) (dto.UserRegisterResponse, error) {
	hashPassword, err := util.HashPassword(req.Password)
	id := uuid.New().String()
	if err != nil {
		return dto.UserRegisterResponse{}, helper.ErrRegisterUser
	}
	user := domain.User{
		ID:       id,
		Username: req.Username,
		Password: hashPassword,
		Email:    req.Email,
	}

	err = u.userRepository.Insert(ctx, &user)
	if err != nil {
		return dto.UserRegisterResponse{}, helper.ErrRegisterUser
	}
	return dto.UserRegisterResponse{
		Status:  true,
		Message: "success register",
	}, nil
}
