package storage

import (
	"database/sql"
	"fmt"
	invoiceitem "go-mysql/pkg/InvoiceItem"
)

const (
	mySQLMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
)

type MySQLItem struct {
	db *sql.DB
}

// CreateTx implements invoiceitem.Storage.
func (*MySQLItem) CreateTx(*sql.Tx, uint, invoiceitem.Models) error {
	panic("unimplemented")
}

func NewMySQLItem(db *sql.DB) *MySQLItem {
	return &MySQLItem{db}
}

func (p *MySQLItem) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de Item ejecutada correctamente")
	return nil
}
