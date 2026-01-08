package dto

import "time"

type InvoiceDetail struct {
	ID        uint64 `json:"id"`
	CompanyID uint64 `json:"company_id"`
	Company   string `json:"companyUsecase"`

	Status string `json:"status"`

	Total int64         `json:"total"`
	Items []InvoiceItem `gorm:"-",json:"items"`

	Approver *string `json:"approver"`

	PaidAt        *time.Time `json:"paid_at,omitempty"`
	PaymentMethod *string    `json:"payment_method,omitempty"`
	PaymentRef    *string    `json:"payment_ref,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InvoiceItem struct {
	Name     string `json:"name"`
	Qty      int64  `json:"qty"`
	Price    int64  `json:"price"`
	SubTotal int64  `json:"sub_total"`
}

type InvoiceList struct {
	ID          uint64 `json:"id"`
	CompanyName string `json:"company_name"`
	Total       int64  `json:"total"`
	Status      string `json:"status"`
}
type InvoiceDashboard struct {
	TotalInvoice int64 `json:"total_invoice"`
	PaidCount    int64 `json:"paid_count"`
	UnpaidCount  int64 `json:"unpaid_count"`
	TotalAmount  int64 `json:"total_amount"`
}

type DashboardRow struct {
	Status string
	Total  int64
}

//func NewInvoiceDashboard(CommpanyID uint64, rows []DashboardRow) *InvoiceDashboard {
//	d := &InvoiceDashboard{
//		CompanyID: CommpanyID,
//	}
//	for _, row := range rows {
//		switch row.Status {
//		case "draft":
//			d.TotalDraft = row.Total
//		case "submited":
//			d.TotalSubmited = row.Total
//		case "approved":
//			d.TotalApproved = row.Total
//		case "paid":
//			d.TotalPaid = row.Total
//		}
//		d.GrandTotal = row.Total
//	}
//	return d
//}

type InvoiceListItem struct {
	ID        uint64    `json:"id"`
	InvoiceNo string    `json:"invoice_no"`
	Total     int64     `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
