package company_query

import "invoice-payment-system/model"

type CompanyReadRepoInterface interface {
	FindByID(id uint64) (*model.Company, error)
	List() ([]model.Company, error)
}
