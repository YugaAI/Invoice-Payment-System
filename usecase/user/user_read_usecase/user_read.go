package user_read_usecase

import (
	"errors"
	"invoice-payment-system/model"

	"golang.org/x/crypto/bcrypt"
)

type UserWriteRepo interface {
	GetByUsername(email string) (*model.User, error)
}

type UserReadUsecase struct {
	Repo UserWriteRepo
}

func NewLoginUsecase(Repo UserWriteRepo) *UserReadUsecase {
	return &UserReadUsecase{
		Repo: Repo,
	}
}

func (u *UserReadUsecase) Login(email, password string) (*model.User, error) {
	user, err := u.Repo.GetByUsername(email)
	if err != nil {
		return nil, errors.New("username not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("password wrong")
	}
	return user, nil
}
