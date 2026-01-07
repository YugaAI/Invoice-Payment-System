package invoice_read

import (
	"context"
	"encoding/json"
	"fmt"
	"invoice-payment-system/dto"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type InvoiceReadRepo struct {
	DB    *gorm.DB
	Redis *redis.Client
	ctx   context.Context
}

func NewInvoiceReadRepo(db *gorm.DB, Redis *redis.Client, ctx context.Context) *InvoiceReadRepo {
	return &InvoiceReadRepo{
		DB:    db,
		Redis: Redis,
		ctx:   ctx,
	}
}

func (r *InvoiceReadRepo) FindDetailByID(id uint64) (*dto.InvoiceDetail, error) {
	var invoice dto.InvoiceDetail
	key := fmt.Sprintf("invoice:detail%d", id)

	if r.Redis != nil {
		if cached, err := r.Redis.Get(r.ctx, key).Result(); err == nil {
			var dto dto.InvoiceDetail
			if json.Unmarshal([]byte(cached), &dto) == nil {
				return &dto, nil
			}
		}
	}

	err := r.DB.Table("invoices i").
		Select(`
				i.id,
				i.company_id,
				c.name as company,
				i.status,
				i.total,
				i.approver,
				i.paid_at,
				i.payment_method,
				i.payment_ref,
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
	if r.Redis != nil {
		b, _ := json.Marshal(invoice)
		_ = r.Redis.Set(r.ctx, key, b, 2*time.Minute).Err()
	}

	invoice.Items = invoiceItem
	return &invoice, nil
}

func (r *InvoiceReadRepo) List(companyID uint64, page, limit int) ([]dto.InvoiceList, error) {
	offset := (page - 1) * limit
	key := fmt.Sprintf(
		"invoice:list:company:%d:page:%d:limit:%d",
		companyID, page, limit,
	)

	if r.Redis != nil {
		if cached, err := r.Redis.Get(r.ctx, key).Result(); err == nil {
			var dto []dto.InvoiceList
			if json.Unmarshal([]byte(cached), &dto) == nil {
				return dto, nil
			}
		}
	}

	var result []dto.InvoiceList
	err := r.DB.Table("invoices i").
		Select(`
			i.id,
			i.total,
			i.status
		`).
		Where("i.company_id = ?", companyID).
		Order("i.id DESC").
		Limit(limit).
		Offset(offset).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	if r.Redis != nil {
		b, _ := json.Marshal(result)
		_ = r.Redis.Set(r.ctx, key, b, 45*time.Second).Err()
	}

	return result, nil
}

func (r *InvoiceReadRepo) GetDashboard(companyID uint64) (*dto.InvoiceDashboard, error) {
	key := fmt.Sprintf("invoice:dashboard:company:%d", companyID)
	if r.Redis != nil {
		if cached, err := r.Redis.Get(r.ctx, key).Result(); err == nil {
			var dto dto.InvoiceDashboard
			if json.Unmarshal([]byte(cached), &dto) == nil {
				return &dto, nil
			}
		}
	}
	var result dto.InvoiceDashboard
	err := r.DB.Raw(`
		SELECT
			COUNT(*)                     AS total_invoice,
			COUNT(*) FILTER (WHERE status = 'PAID') AS paid_count,
			COUNT(*) FILTER (WHERE status != 'PAID') AS unpaid_count,
			COALESCE(SUM(total), 0)       AS total_amount
		FROM invoices
		WHERE company_id = ?`, companyID).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	if r.Redis != nil {
		b, _ := json.Marshal(result)
		_ = r.Redis.Set(r.ctx, key, b, 20*time.Second).Err()
	}

	return &result, nil
}

//func mapRows(rows []dto.DashboardRow) []dto.DashboardRow {
//	result := make([]dto.DashboardRow, 0, len(rows))
//	for _, r := range rows {
//		result = append(result, dto.DashboardRow{
//			Status: r.Status,
//			Total:  r.Total,
//		})
//	}
//	return result
//}
