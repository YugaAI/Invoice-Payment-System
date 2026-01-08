package company_write_usecase

import (
	"invoice-payment-system/model"

	"gorm.io/gorm"
)

type CompanyWriteRepository interface {
	Create(company *model.Company) error
}

type CompanyWriteUsecase struct {
	DB   *gorm.DB
	Repo CompanyWriteRepository
}

func NewCompanyWriteUsecase(db *gorm.DB, repo CompanyWriteRepository) *CompanyWriteUsecase {
	return &CompanyWriteUsecase{
		DB:   db,
		Repo: repo,
	}
}
