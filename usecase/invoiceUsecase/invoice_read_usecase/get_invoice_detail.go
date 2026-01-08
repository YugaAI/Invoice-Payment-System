package invoice_read_usecase

import (
	"invoice-payment-system/dto"
)

func (h *InvoiceReadUsecase) GetInvoiceByIdExecute(q GetInvoiceDetailQuery) (*dto.InvoiceDetail, error) {
	return h.Repo.FindDetailByID(q.InvoiceID)
}
