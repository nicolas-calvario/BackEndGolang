package storage

import (
	"database/sql"
	"fmt"
	"go-postgres/pkg/product"
	"log"
	"sync"
	"time"

	// ...
	_ "github.com/lib/pq" // paquete del driver para PostgreSQL
)

var (
	db   *sql.DB   // estructura db gestiona un pool de conexiones activas e inactivas
	once sync.Once // estructura Once que permite ejecutar una única vez (Singleton)
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		// El primer argumento de Open es el nombre del driver ("postgres")
		// y el segundo argumento es la cadena de conexión, donde se coloca
		// las credenciales de acceso a la BD
		db, err = sql.Open("postgres", "postgres://calvario:root@localhost:15432/DbGo?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

// Pool retorna una unica instancia de db
func Pool() *sql.DB {
	return db
}

// validacion para guaradar nulos
func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.Id,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}
