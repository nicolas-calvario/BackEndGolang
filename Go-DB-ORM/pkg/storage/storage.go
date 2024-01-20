package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB  // estructura db gestiona un pool de conexiones activas e inactivas
	once sync.Once // estructura Once que permite ejecutar una única vez (Singleton)
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open("postgres://calvario:root@localhost:15432/DbGo?sslmode=disable"))
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}
		fmt.Println("Conectado a postgres")
	})
}

// Pool retorna una unica instancia de db
func DB() *gorm.DB {
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

// func scanRowProduct(s scanner) (*product.Model, error) {
// 	m := &product.Model{}
// 	observationNull := sql.NullString{}
// 	updatedAtNull := sql.NullTime{}

// 	err := s.Scan(
// 		&m.Id,
// 		&m.Name,
// 		&observationNull,
// 		&m.Price,
// 		&m.CreatedAt,
// 		&updatedAtNull,
// 	)
// 	if err != nil {
// 		return &product.Model{}, err
// 	}

// 	m.Observations = observationNull.String
// 	m.UpdatedAt = updatedAtNull.Time

// 	return m, nil
// }

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}
