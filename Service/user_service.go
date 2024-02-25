package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/helper"
	"github.com/midoon/e-wallet-go-app-v1/util"
)

type userService struct {
	userRepository domain.UserRepository
	validate       *validator.Validate
}

func NewUserService(userRepository domain.UserRepository, validator *validator.Validate) domain.UserService {
	return &userService{
		userRepository: userRepository,
		validate:       validator,
	}
}

func (u *userService) Register(ctx context.Context, req dto.UserRegisterRequest) error {
	err := u.validate.Struct(req)
	if err != nil {
		return helper.ErrValidation
	}
	countEmail, err := u.userRepository.CountByEmail(ctx, req.Email)
	if err != nil {
		return helper.ErrRegisterUser
	} else if countEmail != 0 {
		return helper.ErrDuplicateData
	}
	hashPassword, err := util.HashPassword(req.Password)
	id := uuid.New().String()
	if err != nil {
		return helper.ErrRegisterUser
	}
	user := domain.User{
		ID:       id,
		Username: req.Username,
		Password: hashPassword,
		Email:    req.Email,
	}

	err = u.userRepository.Insert(ctx, &user)
	if err != nil {
		return helper.ErrRegisterUser
	}
	return nil
}
