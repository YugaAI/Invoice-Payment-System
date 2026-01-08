package company_write_usecase

import (
	"invoice-payment-system/model"
)

func (uc *CompanyWriteUsecase) CreateExecute(name string) (*model.Company, error) {
	company := &model.Company{
		Name: name,
	}
	if err := uc.Repo.Create(company); err != nil {
		return nil, err
	}
	return company, nil
}
