package company_query

import "invoice-payment-system/model"

type ListCompanyUsecase struct {
	Repo CompanyReadRepoInterface
}

//func NewListCompanyUsecase(repo CompanyReadRepoInterface) *ListCompanyUsecase {
//	return &ListCompanyUsecase{Repo: repo}
//}

func (uc *ListCompanyUsecase) Execute() ([]model.Company, error) {
	return uc.Repo.List()
}
