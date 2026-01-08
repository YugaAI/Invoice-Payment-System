package company_read_usecase

import (
	"invoice-payment-system/model"
)

func (uc *CompanyReadUsecase) GetByIdExecute(id uint64) (*model.Company, error) {
	return uc.Repo.FindByID(id)
}
