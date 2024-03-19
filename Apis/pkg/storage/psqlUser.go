package storage

import (
	"Api-Go/pkg/db"
	"Api-Go/pkg/model"
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
	psqlCreateUser  = `INSERT INTO users(name, created_at) VALUES($1, $2) RETURNING id`
	psqlGetAlluser  = `SELECT id, name, created_at, updated_at FROM users`
	psqlGetuserByID = psqlGetAlluser + " WHERE id = $1"
	psqlUpdateuser  = `UPDATE users SET name = $1, updated_at = $2 WHERE id = $3`
	psqlDeleteuser  = `DELETE FROM users WHERE id = $1`
)

type BdUser struct {
	db *sql.DB
}

func NewBdUser(db *sql.DB) *BdUser {
	return &BdUser{db}
}

func (u BdUser) Create(user *model.User) error {

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
	fmt.Println("Se ha guardado correctamente el usuario")
	return nil
}

func (p *BdUser) GetAll() (model.Users, error) {
	stmt, err := p.db.Prepare(psqlGetAlluser)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Users, 0)
	for rows.Next() {
		m := &model.User{}
		updatedAtNull := sql.NullTime{}

		err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.CreatedAt,
			&updatedAtNull,
		)
		if err != nil {
			return nil, err
		}
		m.UpdatedAt = updatedAtNull.Time
		ms = append(ms, *m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

func (p *BdUser) GetById(id int) (*model.User, error) {
	stmt, err := p.db.Prepare(psqlGetuserByID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return db.ScanRowUser(stmt.QueryRow(id))
}

func (p *BdUser) Update(id int, m *model.User) error {
	stmt, err := p.db.Prepare(psqlUpdateuser)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		db.TimeToNull(m.UpdatedAt), // control de nulos (helper en el archivo *storage.go*)
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no existe el producto con id: %d", m.ID)
	}

	fmt.Println("se actualizó el producto correctamente")
	return nil
}

func (p *BdUser) Delete(id int) error {
	stmt, err := p.db.Prepare(psqlDeleteuser)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("se eliminó el producto correctamente")
	return nil
}
