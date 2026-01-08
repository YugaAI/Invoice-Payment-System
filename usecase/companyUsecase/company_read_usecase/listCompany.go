package company_read_usecase

import (
	"invoice-payment-system/model"
)

//func NewListCompanyUsecase(repo CompanyReadRepoInterface) *ListCompanyUsecase {
//	return &ListCompanyUsecase{Repo: repo}
//}

func (uc *CompanyReadUsecase) GetListExecute() ([]model.Company, error) {
	return uc.Repo.List()
}
