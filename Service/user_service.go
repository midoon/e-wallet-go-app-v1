package service

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/helper"
	"github.com/midoon/e-wallet-go-app-v1/internal/config"
	"github.com/midoon/e-wallet-go-app-v1/util"
)

type userService struct {
	userRepository domain.UserRepository
	validate       *validator.Validate
	config         *config.Config
}

// pada provider mengembalikan interface supaya ada error jika terdapat function yang belum diimplement. dan tikad menggunakan pointer pada interface return value, tetapi varibale returnnya harus pointer
func NewUserService(userRepository domain.UserRepository, validator *validator.Validate, config *config.Config) domain.UserService {
	return &userService{
		userRepository: userRepository,
		validate:       validator,
		config:         config,
	}
}

func (u *userService) Register(ctx context.Context, req dto.UserRegisterRequest) error {
	err := u.validate.Struct(req)
	if err != nil {
		return helper.ErrValidation
	}
	countEmail, err := u.userRepository.CountByEmail(ctx, req.Email)
	if err != nil {
		return err
	} else if countEmail != 0 {
		return helper.ErrDuplicateData
	}
	hashPassword, err := util.HashPassword(req.Password)
	id := uuid.New().String()
	if err != nil {
		return err
	}
	user := domain.User{
		ID:       id,
		Username: req.Username,
		Password: hashPassword,
		Email:    req.Email,
	}

	err = u.userRepository.Insert(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) Login(ctx context.Context, req dto.LoginRequest) error {
	err := u.validate.Struct(req)
	if err != nil {
		return helper.ErrValidation
	}
	user, err := u.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return helper.ErrEmailOrPaswordWrong
	}

	if ok := util.CheckPasswordHash(req.Password, user.Password); !ok {
		return helper.ErrEmailOrPaswordWrong
	}

	fmt.Println(user)
	return nil
}
