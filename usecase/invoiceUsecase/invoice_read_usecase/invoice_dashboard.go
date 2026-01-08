package invoice_read_usecase

import (
	"invoice-payment-system/dto"
)

func (h *InvoiceReadUsecase) InvoiceDasboardExecute(q InvoiceDashboardQuery) (*dto.InvoiceDashboard, error) {
	return h.Repo.GetDashboard(q.CompanyID)
}
