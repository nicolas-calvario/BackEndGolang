package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func Pool() *sql.DB {
	return db
}

func NewMysqlDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:root@tcp(localhost:33061)/DbGo")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("conectado a Mysql")
	})
}
