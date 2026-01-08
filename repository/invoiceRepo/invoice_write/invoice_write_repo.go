package invoice_write

import (
	"errors"
	"invoice-payment-system/domain"
	model2 "invoice-payment-system/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InvoiceWriteRepo struct {
	DB *gorm.DB
}

func NewInvoiceWriteRepo(db *gorm.DB) *InvoiceWriteRepo {
	return &InvoiceWriteRepo{
		DB: db,
	}
}

func (i *InvoiceWriteRepo) Create(invoice *domain.Invoice) error {
	items := make([]model2.Item, 0, len(invoice.Items))

	for _, it := range invoice.Items {
		items = append(items, model2.Item{
			Name:     it.Name,
			Qty:      it.Qty,
			Price:    it.Price,
			Subtotal: it.Qty * it.Price,
		})
	}

	model := model2.Invoices{
		ID:        invoice.ID,
		CompanyID: invoice.CompanyID,
		Total:     invoice.Total,
		Status:    string(invoice.Status),
		Items:     items,
	}

	if err := i.DB.Create(&model).Error; err != nil {
		return err
	}

	invoice.ID = model.ID
	return nil
}

func (i *InvoiceWriteRepo) FindByID(id uint64) (*domain.Invoice, error) {
	var model model2.Invoices
	err := i.DB.Clauses(clause.Locking{Strength: "UPDATE"}).First(&model, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &domain.Invoice{
		ID:        model.ID,
		CompanyID: model.CompanyID,
		Total:     model.Total,
		Status:    domain.InvoiceStatus(model.Status),
	}, nil
}

func (r *InvoiceWriteRepo) SaveSubmit(invoice *domain.Invoice) error {
	return r.DB.
		Model(&model2.Invoices{}).
		Where("id = ?", invoice.ID).
		Updates(map[string]interface{}{
			"status":     invoice.Status,
			"updated_at": time.Now(),
		}).Error
}

func (r *InvoiceWriteRepo) SaveApprove(invoice *domain.Invoice) error {
	return r.DB.
		Model(&model2.Invoices{}).
		Where("id = ?", invoice.ID).
		Updates(map[string]interface{}{
			"status":     invoice.Status,
			"approver":   invoice.ApproverBy,
			"updated_at": time.Now(),
		}).Error
}

func (r *InvoiceWriteRepo) SavePayment(invoice *domain.Invoice) error {
	return r.DB.
		Model(&model2.Invoices{}).
		Where("id = ?", invoice.ID).
		Updates(map[string]interface{}{
			"status":         invoice.Status,
			"approver":       invoice.ApproverBy,
			"paid_at":        invoice.PaidAt,
			"payment_ref":    invoice.PaymentRef,
			"payment_method": invoice.PaymentMethod,
			"updated_at":     time.Now(),
		}).Error
}
