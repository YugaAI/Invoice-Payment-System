package invoice_write_usecase

import (
	"gorm.io/gorm"
)

func (h *InvoiceWriteUsecase) SubmitExecute(cmd SubmitInvoiceCommand) error {
	return h.DB.Transaction(func(tx *gorm.DB) error {
		repo := h.Repo
		if txRepo, ok := h.Repo.(interface {
			WithTx(db *gorm.DB) InvoiceWriteRepo
		}); ok {
			repo = txRepo.WithTx(tx)
		}

		invoices, err := repo.FindByID(cmd.InvoiceID)
		if err != nil {
			return err
		}

		if err := invoices.Submit(); err != nil {
			return err
		}
		return repo.SaveSubmit(invoices)
	})
}
