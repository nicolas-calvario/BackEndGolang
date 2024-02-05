package main

import (
	"Api-Go/pkg/db"
	"Api-Go/pkg/handler"
	"Api-Go/pkg/storage"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Creacion y manejo de apis con base de datos de postgres")
	db.NewDbPSQL()

	mux := http.NewServeMux()

	handler.RouterUser(mux, storage.NewBdUser(db.Pool()))
	http.ListenAndServe(":8080", mux)

}
