package invoice_command

import (
	"gorm.io/gorm"
)

type SubmiteInvoiceUsecase struct {
	DB   *gorm.DB
	Repo InvoiceWriteRepoInterface
}

func (h *SubmiteInvoiceUsecase) Execute(cmd SubmitInvoiceCommand) error {
	return h.DB.Transaction(func(tx *gorm.DB) error {
		repo := h.Repo
		if txRepo, ok := h.Repo.(interface {
			WithTx(db *gorm.DB) InvoiceWriteRepoInterface
		}); ok {
			repo = txRepo.WithTx(tx)
		}

		invoice, err := repo.FindByID(cmd.InvoiceID)
		if err != nil {
			return err
		}

		if err := invoice.Submit(); err != nil {
			return err
		}
		return repo.SaveSubmit(invoice)
	})
}
