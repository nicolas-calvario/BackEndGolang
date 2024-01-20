package storage

import (
	"database/sql"
	"fmt"
	"go-mysql/pkg/product"
)

const (
	mySQLMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	mySQLCreateProduct = `INSERT INTO products(name, observations, price, created_at) VALUES(?, ?, ?, ?)`
)

type MySQLProduct struct {
	db *sql.DB
}

func NewMySQLProduct(db *sql.DB) *MySQLProduct {
	return &MySQLProduct{db}
}

func (p *MySQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("migración de producto ejecutada correctamente")
	return nil
}

func (p *MySQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(mySQLCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		m.Name,
		(m.Observations),
		m.Price,
		m.CreatedAt,
	) // aqui no hay método Scan (en Postgres) para la recuperación del id
	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // recuperación del id
	if err != nil {
		return err
	}

	m.Id = uint(id)

	fmt.Printf("se creo el producto correctamente con ID: %d\n", m.Id)
	return nil
}
