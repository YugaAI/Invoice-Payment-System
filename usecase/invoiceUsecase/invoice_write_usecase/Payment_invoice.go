package invoice_write_usecase

import (
	"time"

	"gorm.io/gorm"
)

func (h *InvoiceWriteUsecase) PaymentExecute(cmd PayInvoiceCommand) error {
	if cmd.PaidAt.IsZero() {
		cmd.PaidAt = time.Now()
	}

	return h.DB.Transaction(func(tx *gorm.DB) error {
		repo := h.Repo
		if txRepo, ok := h.Repo.(interface {
			WithTx(*gorm.DB) InvoiceWriteRepo
		}); ok {
			repo = txRepo.WithTx(tx)
		}
		invoices, err := repo.FindByID(cmd.InvoiceID)
		if err != nil {
			return err
		}

		if err := invoices.Pay(cmd.PaidAt, cmd.Method, cmd.RefNo); err != nil {
			return err
		}

		return repo.SavePayment(invoices)
	})
}
