package invoice_write_usecase

import (
	"invoice-payment-system/domain"

	"gorm.io/gorm"
)

func (u *InvoiceWriteUsecase) CreateExecute(cmd CreateInvoiceCommand) (uint64, error) {
	items := make([]domain.InvoiceItem, 0, len(cmd.Items))
	for _, it := range cmd.Items {
		items = append(items, domain.InvoiceItem{
			Name:  it.Name,
			Qty:   it.Qty,
			Price: it.Price,
		})
	}

	invoices, err := domain.NewInvoice(cmd.CompanyID, items)
	if err != nil {
		return 0, err
	}

	err = u.DB.Transaction(func(tx *gorm.DB) error {
		repo := u.Repo
		if txRepo, ok := repo.(interface {
			WithTX(*gorm.DB) InvoiceWriteRepo
		}); ok {
			repo = txRepo.WithTX(tx)
		}
		return repo.Create(invoices)
	})

	if err != nil {
		return 0, err
	}

	return invoices.ID, nil
}
