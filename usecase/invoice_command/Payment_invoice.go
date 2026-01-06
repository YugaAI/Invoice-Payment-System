package invoice_command

import (
	"time"

	"gorm.io/gorm"
)

type PayInvoiceUsecase struct {
	DB   *gorm.DB
	Repo InvoiceWriteRepoInterface
}

func (h *PayInvoiceUsecase) Execute(cmd PayInvoiceCommand) error {
	if cmd.PaidAt.IsZero() {
		cmd.PaidAt = time.Now()
	}

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

		if err := invoice.Pay(cmd.PaidAt, cmd.Method, cmd.RefNo); err != nil {
			return err
		}

		return repo.SavePayment(invoice)
	})
}
