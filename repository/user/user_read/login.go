package user_read

import (
	"context"
	"invoice-payment-system/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Login struct {
	DB          *gorm.DB
	RedisClient *redis.Client
	Ctx         context.Context
}

func NewLogin(db *gorm.DB, redisClient *redis.Client, ctx context.Context) *Login {
	return &Login{
		DB:          db,
		RedisClient: redisClient,
		Ctx:         ctx,
	}
}

func (r *Login) GetByUsername(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", email).First(&user).Error
	return &user, err
}
