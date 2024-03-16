package repository

import (
	"context"
	"log"

	"github.com/midoon/e-wallet-go-app-v1/domain"
	"github.com/midoon/e-wallet-go-app-v1/dto"
	"github.com/midoon/e-wallet-go-app-v1/util"
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

// data yang dikirimkan sebagai query harus berupa pointer
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
		log.Println(err)
		return err
	}
	return nil
}

func (u *userRepository) Resgiter(ctx context.Context, req dto.UserRegisterRequest) error {
	err := u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		hashPassword, err := util.HashPassword(req.Password)
		if err != nil {
			return err
		}

		user := domain.User{
			Username: req.Username,
			Password: hashPassword,
			Email:    req.Email,
		}

		err = tx.Create(&user).Error
		if err != nil {
			return err
		}
		var registeredUser domain.User
		err = tx.Where("email = ?", req.Email).Take(&registeredUser).Error
		if err != nil {
			return err
		}

		hashPin, err := util.HashPassword(req.Pin)
		if err != nil {
			return err
		}
		account := domain.Account{
			AccountNumber: req.AccountNumber,
			Balance:       0,
			Pin:           hashPin,
			UserId:        registeredUser.ID,
		}
		err = tx.Create(&account).Error
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		log.Println(err)
	}
	return err
}
