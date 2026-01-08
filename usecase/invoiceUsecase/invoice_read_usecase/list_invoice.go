package invoice_read_usecase

import (
	"invoice-payment-system/dto"
)

func (u *InvoiceReadUsecase) GetListInvoiceExecute(companyID uint64, page, limit int) ([]dto.InvoiceList, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	return u.Repo.List(companyID, page, limit)
}
