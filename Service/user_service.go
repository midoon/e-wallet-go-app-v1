package service

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/helper"
	"github.com/midoon/e-wallet-go-app-v1/internal/config"
	"github.com/midoon/e-wallet-go-app-v1/util"
)

type userService struct {
	userRepository  domain.UserRepository
	tokenRepository domain.TokenRepository
	validate        *validator.Validate
	config          *config.Config
}

// pada provider mengembalikan interface supaya ada error jika terdapat function yang belum diimplement. dan tikad menggunakan pointer pada interface return value, tetapi varibale returnnya harus pointer
func NewUserService(userRepository domain.UserRepository, tokenRepository domain.TokenRepository, validator *validator.Validate, config *config.Config) domain.UserService {
	return &userService{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		validate:        validator,
		config:          config,
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
	if err != nil {
		return err
	}
	id := uuid.New().String()
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

func (u *userService) Login(ctx context.Context, req dto.LoginRequest) (dto.TokenData, error) {
	err := u.validate.Struct(req)
	if err != nil {
		return dto.TokenData{}, helper.ErrValidation
	}
	user, err := u.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.TokenData{}, helper.ErrEmailOrPaswordWrong
	}

	if ok := util.CheckPasswordHash(req.Password, user.Password); !ok {
		return dto.TokenData{}, helper.ErrEmailOrPaswordWrong
	}

	countToken, err := u.tokenRepository.CountByUserId(ctx, user.ID)

	if err != nil {
		return dto.TokenData{}, err
	}

	if countToken != 0 {
		err := u.tokenRepository.Delete(ctx, user.ID)
		if err != nil {
			return dto.TokenData{}, err
		}
	}

	// generate access token
	atExpTime := time.Now().Add(time.Hour * 3)
	ATClaim := util.NewJwtClaim(user.ID, user.Email, u.config.JWT.Issuer, atExpTime)
	aToken, err := ATClaim.SignToken(u.config.JWT.Key)
	if err != nil {
		return dto.TokenData{}, err
	}

	// generate refresh token
	rtExpTime := time.Now().Add(time.Hour * 24 * 3)
	RTClaim := util.NewJwtClaim(user.ID, user.Email, u.config.JWT.Issuer, rtExpTime)
	rToken, err := RTClaim.SignToken(u.config.JWT.Key)
	if err != nil {
		return dto.TokenData{}, err
	}

	// save refresh token to db
	token := domain.Token{
		UserId:       user.ID,
		RefreshToken: rToken,
	}

	err = u.tokenRepository.Insert(ctx, &token)
	if err != nil {
		return dto.TokenData{}, err
	}

	tokenData := dto.TokenData{
		UserId:       user.ID,
		AccessToken:  aToken,
		RefreshToken: rToken,
	}

	return tokenData, nil
}

func (u *userService) Logout(ctx context.Context, userId string) error {
	countUser, err := u.tokenRepository.CountByUserId(ctx, userId)
	if err != nil {
		return err
	}
	if countUser == 0 {
		return helper.ErrAccessDenied
	}
	err = u.tokenRepository.Delete(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) Refresh(ctx context.Context, req dto.RefreshRequest) (dto.RefreshData, error) {
	err := u.validate.Struct(req)
	if err != nil {
		return dto.RefreshData{}, nil
	}

	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, helper.ErrJwtValidation
		} else if method != jwt.SigningMethodHS256 {
			return nil, helper.ErrJwtValidation
		}

		return []byte(u.config.JWT.Key), nil
	})

	if err != nil {
		return dto.RefreshData{}, err
	}
	claims := token.Claims.(jwt.MapClaims)

	// check is user logged
	countUser, err := u.tokenRepository.CountByUserId(ctx, claims["Id"].(string))
	if err != nil {
		return dto.RefreshData{}, err
	}
	if countUser == 0 {
		return dto.RefreshData{}, helper.ErrAccessDenied
	}

	// generate new access token
	atExpTime := time.Now().Add(time.Hour * 3)
	ATClaim := util.NewJwtClaim(claims["Id"].(string), claims["Email"].(string), u.config.JWT.Issuer, atExpTime)
	aToken, err := ATClaim.SignToken(u.config.JWT.Key)
	if err != nil {
		return dto.RefreshData{}, err
	}

	refreshData := dto.RefreshData{
		AccessToken: aToken,
	}
	return refreshData, nil
}
