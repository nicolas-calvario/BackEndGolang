package main

import (
	"Api-Go/authorization"
	"Api-Go/pkg/db"
	"Api-Go/pkg/handler"
	"Api-Go/pkg/storage"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Creacion y manejo de apis con base de datos de postgres")
	db.NewDbPSQL()
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se puedo cargar los certificados: %v", err)
	}

	mux := http.NewServeMux()

	handler.RouterUser(mux, storage.NewBdUser(db.Pool()))
	handler.RouterLogin(mux, storage.NewBdUser(db.Pool()))

	http.ListenAndServe(":8080", mux)

}
