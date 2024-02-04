package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateUser = `CREATE TABLE IF NOT EXISTS Users(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT user_id_pk PRIMARY KEY (id) 
	)`
	// psqlCreateProduct = `INSERT INTO products(name, observations, price, created_at) VALUES($1, $2, $3, $4) RETURNING id`
	// psqlGetAllProduct = `SELECT id, name, observations, price,
	// created_at, updated_at
	// FROM products`
	// psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"
	// psqlUpdateProduct  = `UPDATE products SET name = $1, observations = $2,
	// price = $3, updated_at = $4 WHERE id = $5`
	// psqlDeleteProduct = `DELETE FROM products WHERE id = $1`
)

type BdUser struct {
	db *sql.DB
}

func NewBdUser(db *sql.DB) *BdUser {
	return &BdUser{db}
}

func (u BdUser) Migrate() error {
	stmt, err := u.db.Prepare(psqlMigrateUser)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Migración de la tabla Usuario ejecutada correctamente")
	return nil
}
