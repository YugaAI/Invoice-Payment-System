package model

import "time"

type Invoices struct {
	ID        uint64 `gorm:"primaryKey"`
	CompanyID uint64 `gorm:"not null;index"`
	Total     int64  `gorm:"not null"`
	Status    string `gorm:"type:varchar(20);not null"`

	Items []Item `gorm:"foreignKey:InvoiceID;constraint:OnDelete:CASCADE;"`

	Approver string `gorm:"type:varchar(20);not null"`

	PaidAt time.Time `gorm:"not null"`

	PaymentMethod string `gorm:"type:varchar(20);not null"`
	PaymentRef    string `gorm:"type:varchar(20);not null"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
type Company struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
}

type Item struct {
	ID        uint64 `gorm:"primaryKey"`
	InvoiceID uint64 `gorm:"not null;index"`
	Name      string
	Qty       int64
	Price     int64
	Subtotal  int64
	CreatedAt time.Time
}
