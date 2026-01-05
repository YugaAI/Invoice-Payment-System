package company_command

import "invoice-payment-system/model"

type CompanyWriteRepoInterface interface {
	Create(company *model.Company) error
}
