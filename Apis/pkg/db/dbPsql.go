package db

import (
	"Api-Go/pkg/model"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewDbPSQL() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://calvario:root@localhost:15432/DbGo?sslmode=disable")
		if err != nil {
			log.Fatalf("Error al conectar con base de datos:%v ", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("No existe comunicacion con la base de datos")
		}
		fmt.Println("Conectado a base de datos")
	})
}

func Pool() *sql.DB {
	return db
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func ScanRowUser(s scanner) (*model.User, error) {
	m := &model.User{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return &model.User{}, err
	}

	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}

func TimeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}
