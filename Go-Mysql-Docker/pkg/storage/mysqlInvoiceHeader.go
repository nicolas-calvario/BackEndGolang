package storage

import (
	"database/sql"
	"fmt"
	invoiceheader "go-mysql/pkg/invoiceHeader"
)

const (
	mySQLMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
)

type MySQLHeader struct {
	db *sql.DB
}

// CreateTx implements invoiceheader.Storage.
func (*MySQLHeader) CreateTx(*sql.Tx, *invoiceheader.Model) error {
	panic("unimplemented")
}

func NewMySQLHeader(db *sql.DB) *MySQLHeader {
	return &MySQLHeader{db}
}

func (p *MySQLHeader) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de Header ejecutada correctamente")
	return nil
}
