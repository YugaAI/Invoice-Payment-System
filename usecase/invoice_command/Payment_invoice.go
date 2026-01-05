package invoice_command

import (
	"errors"
	"invoice-payment-system/domain"
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
			withTx(*gorm.DB) InvoiceWriteRepoInterface
		}); ok {
			repo = txRepo.withTx(tx)
		}
		invoice, err := repo.FindByID(cmd.InvoiceID)
		if err != nil {
			return err
		}
		if invoice.Status != domain.Approved {
			return errors.New("only approved invoices can be paid")
		}
		invoice.Pay(cmd.PaidAt, cmd.Method, cmd.RefNo)

		return repo.Save(invoice)
	})
}
