package storage

import (
	"database/sql"
	"fmt"
	"go-postgres/pkg/product"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price FLOAT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id) 
	)`
	psqlCreateProduct = `INSERT INTO products(name, observations, price, created_at) VALUES($1, $2, $3, $4) RETURNING id`
	psqlGetAllProduct = `SELECT id, name, observations, price, 
	created_at, updated_at
	FROM products`
	psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"
)

type PsqlProduct struct {
	db *sql.DB
}

func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de la tabla producto ejecutada correctamente")
	return nil
}

func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.Id)
	if err != nil {
		return err
	}
	fmt.Println("Se ha guardado correctamente el producto")
	return nil
}

func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m := &product.Model{}
		observationNull := sql.NullString{}
		updatedAtNull := sql.NullTime{}

		err := rows.Scan(
			&m.Id,
			&m.Name,
			&observationNull,
			&m.Price,
			&m.CreatedAt,
			&updatedAtNull,
		)
		if err != nil {
			return nil, err
		}
		m.Observations = observationNull.String
		m.UpdatedAt = updatedAtNull.Time
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

func (p *PsqlProduct) GetById(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return scanRowProduct(stmt.QueryRow(id))
}
