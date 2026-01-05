package invoice_command

import (
	"errors"
	"invoice-payment-system/domain"

	"gorm.io/gorm"
)

type ApproveInvoiceUsecase struct {
	DB   *gorm.DB
	Repo InvoiceWriteRepoInterface
}

func (h *ApproveInvoiceUsecase) Execute(cmd ApproveInvoiceCommand) error {
	return h.DB.Transaction(func(tx *gorm.DB) error {
		repo := h.Repo
		if txRepo, ok := h.Repo.(interface {
			withTx(*gorm.DB) InvoiceWriteRepoInterface
		}); ok {
			repo = txRepo.withTx(tx)
		}
		invoice, err := repo.FindByID(cmd.InvoiceID)
		if err != nil {
			return err
		}

		if invoice.Status != domain.Submitted {
			return errors.New("only submitted invoices can be approved")
		}

		invoice.Approve(cmd.Approver)
		return repo.Save(invoice)
	})
}
