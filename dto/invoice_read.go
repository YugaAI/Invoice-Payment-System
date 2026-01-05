package dto

import "time"

type InvoiceDetail struct {
	ID        uint64 `json:"id"`
	CompanyID uint64 `json:"company_id"`
	Company   string `json:"company"`

	Status string `json:"status"`

	Total int64         `json:"total"`
	Items []InvoiceItem `gorm:"-",json:"items"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InvoiceItem struct {
	Name     string `json:"name"`
	Qty      int64  `json:"qty"`
	Price    int64  `json:"price"`
	SubTotal int64  `json:"sub_total"`
}

type InvoiceDashboard struct {
	CompanyID uint64 `json:"company_id"`

	TotalDraft    int64 `json:"total_draft"`
	TotalSubmited int64 `json:"total_submited"`
	TotalApproved int64 `json:"total_approved"`
	TotalPaid     int64 `json:"total_paid"`

	GrandTotal int64 `json:"grand_total"`
}

type DashboardRow struct {
	Status string
	Total  int64
}

func NewInvoiceDashboard(CommpanyID uint64, rows []DashboardRow) *InvoiceDashboard {
	d := &InvoiceDashboard{
		CompanyID: CommpanyID,
	}
	for _, row := range rows {
		switch row.Status {
		case "draft":
			d.TotalDraft = row.Total
		case "submited":
			d.TotalSubmited = row.Total
		case "approved":
			d.TotalApproved = row.Total
		case "paid":
			d.TotalPaid = row.Total
		}
		d.GrandTotal = row.Total
	}
	return d
}

type InvoiceListItem struct {
	ID        uint64    `json:"id"`
	InvoiceNo string    `json:"invoice_no"`
	Total     int64     `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
