package invoice_query

import "invoice-payment-system/dto"

type GetInvoiceDetailQuery struct {
	InvoiceID uint64
}

type ListInvoiceQuery struct {
	CompanyID uint64
	Status    *string
	Limit     int
	Offset    int
}

type ListInvoiceResult struct {
	Data  []dto.InvoiceListItem
	Total int64
}

type InvoiceDashboardQuery struct {
	CompanyID uint64
}
