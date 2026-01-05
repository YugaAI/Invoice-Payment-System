package invoice_command

import (
	"errors"
	"invoice-payment-system/domain"

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
			withTx(db *gorm.DB) InvoiceWriteRepoInterface
		}); ok {
			repo = txRepo.withTx(tx)
		}

		invoice, err := repo.FindByID(cmd.InvoiceID)
		if err != nil {
			return err
		}

		if invoice.Status != domain.Draft {
			return errors.New("only draft invoice can be submitted")
		}

		if err := invoice.Submit(); err != nil {
			return err
		}
		return repo.Save(invoice)
	})
}
