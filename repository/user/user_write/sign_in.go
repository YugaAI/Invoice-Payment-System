package user_write

import (
	"invoice-payment-system/model"

	"gorm.io/gorm"
)

type SignIn struct {
	db *gorm.DB
}

func NewSignIn(db *gorm.DB) *SignIn {
	return &SignIn{
		db: db,
	}
}

func (r *SignIn) Create(user *model.User) error {
	return r.db.Create(user).Error
}
