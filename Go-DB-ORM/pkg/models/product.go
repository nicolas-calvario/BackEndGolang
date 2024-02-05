package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string  `gorm: "type: varchar(256); not null"`
	Observations *string `gorm: "type: varchar(256)"`
	Price        float64 `gorm: "not null"`
	invoiceItems []InvoiceItem
}

type InvoiceHeader struct {
	gorm.Model
	Client       string `gorm: "type: varchar(256); not null"`
	invoiceItems []InvoiceItem
}

type InvoiceItem struct {
	gorm.Model
	InvoiceHeaderID string
	ProductID       uint
}
