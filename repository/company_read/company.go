package company_read

import (
	"invoice-payment-system/model"

	"gorm.io/gorm"
)

type CompanyReadRepo struct {
	db *gorm.DB
}

func NewCompanyReadRepo(db *gorm.DB) *CompanyReadRepo {
	return &CompanyReadRepo{db: db}
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
