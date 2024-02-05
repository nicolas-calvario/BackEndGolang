package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewDbPSQL() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://calvario:root@localhost:15432/GoApis?sslmode=disable")
		if err != nil {
			log.Fatalf("Error al conectar con base de datos:%v ", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("No existe comunicacion con la base de datos")
		}
		fmt.Println("Conectado a")
	})
}

func Pool() *sql.DB {
	return db
}
