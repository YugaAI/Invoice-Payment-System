package invoice_command

import (
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
			WithTx(*gorm.DB) InvoiceWriteRepoInterface
		}); ok {
			repo = txRepo.WithTx(tx)
		}
		invoice, err := repo.FindByID(cmd.InvoiceID)
		if err != nil {
			return err
		}

		if err := invoice.Approve(cmd.Approver); err != nil {
			return err
		}
		return repo.SaveApprove(invoice)
	})
}
