package invoice_command

import "time"

type CreateInvoiceItem struct {
	Name  string `json:"name" binding:"required"`
	Qty   int64  `json:"quantity" binding:"required,gt=0"`
	Price int64  `json:"price" binding:"required,gt=0"`
}

type CreateInvoiceCommand struct {
	CompanyID uint64              `json:"company_id" binding:"required"`
	Items     []CreateInvoiceItem `json:"items" binding:"required,dive"`
}

type SubmitInvoiceCommand struct {
	InvoiceID uint64
}

type ApproveInvoiceCommand struct {
	InvoiceID uint64
	Approver  string
}

type PayInvoiceCommand struct {
	InvoiceID uint64
	PaidAt    time.Time
	Method    string
	RefNo     string
}
