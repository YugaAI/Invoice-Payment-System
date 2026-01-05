package domain

import (
	"errors"
	"time"
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
		return errors.New("only draft can be submitted")
	}
	i.Status = Submitted
	return nil
}

func (i *Invoice) Approve(approver string) error {
	if i.Status != Submitted {
		return errors.New("only approved can be approved")
	}
	i.Status = Approved
	i.ApproverBy = approver
	return nil
}
func (i *Invoice) Pay(paidAt time.Time, method, ref string) error {
	if i.Status != Approved {
		return errors.New("only payed can be approved")
	}
	i.Status = Paid
	i.PaidAt = paidAt
	i.PaymentMethod = method
	i.PaymentRef = ref
	return nil
}
