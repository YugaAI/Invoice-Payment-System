package invoice_write_usecase

import (
	"invoice-payment-system/domain"

	"gorm.io/gorm"
)

type InvoiceWriteRepo interface {
	Create(invoice *domain.Invoice) error
	FindByID(id uint64) (*domain.Invoice, error)
	SaveSubmit(invoice *domain.Invoice) error
	SaveApprove(invoice *domain.Invoice) error
	SavePayment(invoice *domain.Invoice) error
}

type InvoiceWriteUsecase struct {
	DB   *gorm.DB
	Repo InvoiceWriteRepo
}

func NewInvoiceWriteUsecase(db *gorm.DB, repo InvoiceWriteRepo) *InvoiceWriteUsecase {
	return &InvoiceWriteUsecase{
		DB:   db,
		Repo: repo,
	}
}
