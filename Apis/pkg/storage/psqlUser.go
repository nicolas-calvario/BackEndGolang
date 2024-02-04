package storage

import (
	user "Api-Go/pkg/User"
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
	psqlCreateUser = `INSERT INTO users(name, created_at) VALUES($1, $2) RETURNING id`
	// psqlGetAlluser = `SELECT id, name, observations, price,
	// created_at, updated_at
	// FROM users`
	// psqlGetuserByID = psqlGetAlluser + " WHERE id = $1"
	// psqlUpdateuser  = `UPDATE users SET name = $1, observations = $2,
	// price = $3, updated_at = $4 WHERE id = $5`
	// psqlDeleteuser = `DELETE FROM users WHERE id = $1`
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

func (u BdUser) Create(user *user.Model) error {

	stmt, err := u.db.Prepare(psqlCreateUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		user.Name,
		user.CreatedAt,
	).Scan(&user.ID)
	if err != nil {
		return err
	}
	fmt.Println("Se ha guardado correctamente el usero")
	return nil
}
