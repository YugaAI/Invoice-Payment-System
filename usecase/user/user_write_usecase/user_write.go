package user_write_usecase

import (
	"invoice-payment-system/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserWriteRepo interface {
	Create(user *model.User) error
}

type UserWriteUsecase struct {
	DB   *gorm.DB
	Repo UserWriteRepo
}

func NewWriteUsecase(DB *gorm.DB, Repo UserWriteRepo) *UserWriteUsecase {
	return &UserWriteUsecase{
		DB:   DB,
		Repo: Repo,
	}
}

func (u *UserWriteUsecase) SignUp(username, email, password, role string) (*model.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Username: username,
		Email:    email,
		Password: string(hashed),
		Role:     role,
	}
	err = u.Repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
