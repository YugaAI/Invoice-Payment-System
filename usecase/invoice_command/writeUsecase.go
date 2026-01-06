package invoice_command

import "invoice-payment-system/domain"

type InvoiceWriteRepoInterface interface {
	Create(invoice *domain.Invoice) error
	FindByID(id uint64) (*domain.Invoice, error)
	SaveSubmit(invoice *domain.Invoice) error
	SaveApprove(invoice *domain.Invoice) error
	SavePayment(invoice *domain.Invoice) error
}

//type WriteUsecase struct {
//	Repo InvoiceWriteRepoInterface
//}
//
//func NewUsecase(repo InvoiceWriteRepoInterface) *WriteUsecase {
//	return &WriteUsecase{
//		Repo: repo,
//	}
//}
