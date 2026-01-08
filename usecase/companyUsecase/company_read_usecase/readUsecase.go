package company_read_usecase

import (
	"invoice-payment-system/model"
)

type CompanyReadRepository interface {
	FindByID(id uint64) (*model.Company, error)
	List() ([]model.Company, error)
}

type CompanyReadUsecase struct {
	Repo CompanyReadRepository
}

func NewCompanyReadUsecase(Repo CompanyReadRepository) *CompanyReadUsecase {
	return &CompanyReadUsecase{
		Repo: Repo,
	}
}
