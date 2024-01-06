package storage

import (
	"database/sql"
	"fmt"
	invoiceithem "go-postgres/pkg/InvoiceItem"
	"go-postgres/pkg/invoice"
	invoiceheader "go-postgres/pkg/invoiceHeader"
)

type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceithem.Storage
}

func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceithem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

func (p *PsqlInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return fmt.Errorf("Header: %w", err)
	}
	fmt.Printf("Factura creada con id: %d \n", m.Header.Id)

	if err := p.storageItems.CreateTx(tx, m.Header.Id, m.Items); err != nil {
		tx.Rollback()
		return fmt.Errorf("Items: %w", err)
	}
	fmt.Printf("items creados: %d \n", len(m.Items))

	return tx.Commit()
}
