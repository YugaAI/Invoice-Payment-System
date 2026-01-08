package invoice_write_usecase

import (
	"gorm.io/gorm"
)

func (h *InvoiceWriteUsecase) ApproveExecute(cmd ApproveInvoiceCommand) error {
	return h.DB.Transaction(func(tx *gorm.DB) error {
		repo := h.Repo
		if txRepo, ok := h.Repo.(interface {
			WithTx(*gorm.DB) InvoiceWriteRepo
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
