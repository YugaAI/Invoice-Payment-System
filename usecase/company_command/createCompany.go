package company_command

import "invoice-payment-system/model"

type CreateCompanyUsecase struct {
	Repo CompanyWriteRepoInterface
}

//func NewCreateCompanyUsecase(repo CompanyWriteRepoInterface) *CreateCompanyUsecase {
//	return &CreateCompanyUsecase{Repo: repo}
//}

func (uc *CreateCompanyUsecase) Execute(name string) (*model.Company, error) {
	company := &model.Company{
		Name: name,
	}
	if err := uc.Repo.Create(company); err != nil {
		return nil, err
	}
	return company, nil
}
