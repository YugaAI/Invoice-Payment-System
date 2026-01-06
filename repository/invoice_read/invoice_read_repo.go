package invoice_read

import (
	"invoice-payment-system/dto"

	"gorm.io/gorm"
)

type InvoiceReadRepo struct {
	DB *gorm.DB
}

func NewInvoiceReadRepo(db *gorm.DB) *InvoiceReadRepo {
	return &InvoiceReadRepo{
		DB: db,
	}
}

func (r *InvoiceReadRepo) FindDetailByID(id uint64) (*dto.InvoiceDetail, error) {
	var invoice dto.InvoiceDetail

	err := r.DB.Table("invoices i").
		Select(`
				i.id,
				i.company_id,
				c.name as company,
				i.status,
				i.total,
				i.created_at,
				i.updated_at
			`).
		Joins("JOIN companies c ON c.id = i.company_id").
		Where("i.id=?", id).
		Scan(&invoice).Error

	if err != nil {
		return nil, err
	}

	var invoiceItem []dto.InvoiceItem
	err = r.DB.Table("items").
		Select(`name, qty, price, (qty*price) as sub_total`).
		Where("invoice_id=?", id).Scan(&invoiceItem).Error

	if err != nil {
		return nil, err
	}
	invoice.Items = invoiceItem
	return &invoice, nil
}

func (r *InvoiceReadRepo) List(companyID uint64, status *string, limit, offset int) ([]dto.InvoiceListItem, int64, error) {
	var (
		items []dto.InvoiceListItem
		total int64
	)
	q := r.DB.Table("invoices").Where("company_id=?", companyID)

	if status != nil {
		q = q.Where("status=?", *status)
	}

	err := q.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	errData := q.Select(`id, total, status, created_at`).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).Scan(&items).Error

	if errData != nil {
		return nil, 0, errData
	}

	return items, total, nil
}

func (r *InvoiceReadRepo) GetDashboard(companyID uint64) (*dto.InvoiceDashboard, error) {
	var rows []dto.DashboardRow

	err := r.DB.Table("invoices").
		Select("status, SUM(total) as total").
		Where("company_id=?", companyID).
		Group("status").
		Scan(&rows).Error

	if err != nil {
		return nil, err
	}

	return dto.NewInvoiceDashboard(companyID, mapRows(rows)), nil
}

func mapRows(rows []dto.DashboardRow) []dto.DashboardRow {
	result := make([]dto.DashboardRow, 0, len(rows))
	for _, r := range rows {
		result = append(result, dto.DashboardRow{
			Status: r.Status,
			Total:  r.Total,
		})
	}
	return result
}
