package invoice_command

import (
	"invoice-payment-system/domain"

	"gorm.io/gorm"
)

type CreateInvoiceUsecase struct {
	DB   *gorm.DB
	Repo InvoiceWriteRepoInterface
}

func (u *CreateInvoiceUsecase) Execute(cmd CreateInvoiceCommand) (uint64, error) {
	items := make([]domain.InvoiceItem, 0, len(cmd.Items))
	for _, it := range cmd.Items {
		items = append(items, domain.InvoiceItem{
			Name:  it.Name,
			Qty:   it.Qty,
			Price: it.Price,
		})
	}

	invoice, err := domain.NewInvoice(cmd.CompanyID, items)
	if err != nil {
		return 0, err
	}

	err = u.DB.Transaction(func(tx *gorm.DB) error {
		repo := u.Repo
		if txRepo, ok := repo.(interface {
			WithTX(*gorm.DB) InvoiceWriteRepoInterface
		}); ok {
			repo = txRepo.WithTX(tx)
		}
		return repo.Create(invoice)
	})

	if err != nil {
		return 0, err
	}

	return invoice.ID, nil
}
