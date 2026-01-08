package invoice_read_usecase

import (
	"invoice-payment-system/dto"
)

type InvoiceReadRepo interface {
	FindDetailByID(id uint64) (*dto.InvoiceDetail, error)
	List(companyID uint64, page, limit int) ([]dto.InvoiceList, error)
	GetDashboard(companyID uint64) (*dto.InvoiceDashboard, error)
}

type InvoiceReadUsecase struct {
	Repo InvoiceReadRepo
}

func NewInvoiceReadUsecase(Repo InvoiceReadRepo) *InvoiceReadUsecase {
	return &InvoiceReadUsecase{
		Repo: Repo,
	}
}
