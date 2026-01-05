package company_query

import "invoice-payment-system/model"

type GetCompanyUsecase struct {
	Repo CompanyReadRepoInterface
}

//func NewGetCompanyUsecase(repo CompanyReadRepoInterface) *GetCompanyUsecase {
//	return &GetCompanyUsecase{Repo: repo}
//}

func (uc *GetCompanyUsecase) Execute(id uint64) (*model.Company, error) {
	return uc.Repo.FindByID(id)
}
