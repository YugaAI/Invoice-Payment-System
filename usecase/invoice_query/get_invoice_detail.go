package invoice_query

import (
	"invoice-payment-system/dto"
)

type GetInvoiceDetailUsecase struct {
	Repo InvoiceReadRepoInterface
}

func (h *GetInvoiceDetailUsecase) Execute(q GetInvoiceDetailQuery) (*dto.InvoiceDetail, error) {
	return h.Repo.FindDetailByID(q.InvoiceID)
}
