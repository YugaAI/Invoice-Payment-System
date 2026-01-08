package company_read

import (
	"context"
	"invoice-payment-system/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CompanyReadRepo struct {
	db    *gorm.DB
	Redis *redis.Client
	ctx   context.Context
}

func NewCompanyReadRepo(db *gorm.DB, Redis *redis.Client, ctx context.Context) *CompanyReadRepo {
	return &CompanyReadRepo{
		db:    db,
		Redis: Redis,
		ctx:   ctx,
	}
}

func (r *CompanyReadRepo) FindByID(id uint64) (*model.Company, error) {
	var company model.Company
	if err := r.db.First(&company, id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}
func (r *CompanyReadRepo) List() ([]model.Company, error) {
	var companies []model.Company
	return companies, r.db.Find(&companies).Error
}
