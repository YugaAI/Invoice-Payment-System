package domain

import (
	"errors"
	"time"
)

var (
	ErrInvoiceEmptyItem    = errors.New("invoiceRepo must have at least one item")
	ErrInvalidInvoiceItem  = errors.New("invalid invoiceRepo item")
	ErrInvoiceNotDraft     = errors.New("only draft invoiceRepo can be submitted")
	ErrInvoiceNotSubmitted = errors.New("only submitted invoiceRepo can be approved")
	ErrInvalidApprover     = errors.New("approver is required")
	ErrInvoiceNotApproved  = errors.New("only approved invoiceRepo can be paid")
	ErrInvoiceAlreadyPaid  = errors.New("invoiceRepo already paid")
	ErrInvalidPayment      = errors.New("invalid payment data")
)

type InvoiceStatus string

const (
	Draft     InvoiceStatus = "DRAFT"
	Submitted InvoiceStatus = "SUBMITTED"
	Approved  InvoiceStatus = "APPROVED"
	Paid      InvoiceStatus = "PAID"
)

type Invoice struct {
	ID         uint64
	CompanyID  uint64
	Total      int64
	Status     InvoiceStatus
	Items      []InvoiceItem
	ApproverBy string

	PaidAt        time.Time
	PaymentMethod string
	PaymentRef    string
}

type InvoiceItem struct {
	Name  string
	Qty   int64
	Price int64
}

func (i *Invoice) Submit() error {
	if i.Status != Draft {
		return ErrInvoiceNotDraft
	}
	i.Status = Submitted
	return nil
}

func (i *Invoice) Approve(approver string) error {
	if i.Status != Submitted {
		return ErrInvoiceNotSubmitted
	}
	if approver == "" {
		return ErrInvalidApprover
	}
	i.Status = Approved
	i.ApproverBy = approver
	return nil
}
func (i *Invoice) Pay(paidAt time.Time, method, ref string) error {
	if i.Status == Paid {
		return ErrInvoiceAlreadyPaid
	}
	if i.Status != Approved {
		return ErrInvoiceNotApproved
	}
	if paidAt.IsZero() || method == "" || ref == "" {
		return ErrInvalidPayment
	}

	i.Status = Paid
	i.PaidAt = paidAt
	i.PaymentMethod = method
	i.PaymentRef = ref
	return nil
}

func NewInvoice(companyID uint64, items []InvoiceItem) (*Invoice, error) {
	if len(items) == 0 {
		return nil, ErrInvoiceEmptyItem
	}

	var total int64
	for _, it := range items {
		if it.Qty <= 0 || it.Price <= 0 {
			return nil, ErrInvalidInvoiceItem
		}
		total += it.Qty * it.Price
	}

	return &Invoice{
		CompanyID: companyID,
		Status:    Draft,
		Total:     total,
		Items:     items,
	}, nil
}
