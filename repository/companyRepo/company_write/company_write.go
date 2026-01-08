package company_write

import (
	"invoice-payment-system/model"

	"gorm.io/gorm"
)

type CompanyWriteRepo struct {
	db *gorm.DB
}

func NewCompanyWriteRepo(db *gorm.DB) *CompanyWriteRepo {
	return &CompanyWriteRepo{db: db}
}

func (r *CompanyWriteRepo) Create(company *model.Company) error {
	return r.db.Create(company).Error
}
