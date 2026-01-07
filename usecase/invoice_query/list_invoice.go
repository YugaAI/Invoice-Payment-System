package invoice_query

import "invoice-payment-system/dto"

type ListInvoiceUsecase struct {
	Repo InvoiceReadRepoInterface
}

func (u *ListInvoiceUsecase) Execute(companyID uint64, page, limit int) ([]dto.InvoiceList, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	return u.Repo.List(companyID, page, limit)
}
