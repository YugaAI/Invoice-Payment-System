package invoice_query

import (
	"invoice-payment-system/dto"
)

type InvoiceDashboardUsecase struct {
	Repo InvoiceReadRepoInterface
}

func (h *InvoiceDashboardUsecase) Execute(q InvoiceDashboardQuery) (*dto.InvoiceDashboard, error) {
	return h.Repo.GetDashboard(q.CompanyID)
}
