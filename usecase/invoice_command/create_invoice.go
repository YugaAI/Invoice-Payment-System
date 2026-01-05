package invoice_command

import (
	"errors"
	"invoice-payment-system/domain"

	"gorm.io/gorm"
)

type CreateInvoiceUsecase struct {
	DB   *gorm.DB
	Repo InvoiceWriteRepoInterface
}

func (h *CreateInvoiceUsecase) Execute(cmd CreateInvoiceCommand) (uint64, error) {
	if len(cmd.Items) == 0 {
		return 0, errors.New("invoice must have at least one item")
	}

	var total int64
	for _, item := range cmd.Items {
		if item.Qty <= 0 || item.Price <= 0 {
			return 0, errors.New("invalid quantity or price")
		}
		total += item.Qty * item.Price

	}

	invoice := &domain.Invoice{
		CompanyID: cmd.CompanyID,
		Total:     total,
		Status:    domain.Draft,
		Items:     make([]domain.InvoiceItem, 0, len(cmd.Items)),
	}
	for _, item := range cmd.Items {
		invoice.Items = append(invoice.Items, domain.InvoiceItem{
			Name:  item.Name,
			Qty:   item.Qty,
			Price: item.Price,
		})
	}

	err := h.DB.Transaction(func(tx *gorm.DB) error {
		repo := h.Repo
		if txRepo, ok := h.Repo.(interface {
			withTX(*gorm.DB) InvoiceWriteRepoInterface
		}); ok {
			repo = txRepo.withTX(tx)
		}

		return repo.Create(invoice)
	})

	if err != nil {
		return 0, err
	}

	return invoice.ID, nil
}
