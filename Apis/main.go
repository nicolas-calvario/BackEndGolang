package main

import (
	user "Api-Go/pkg/User"
	db "Api-Go/pkg/dB"
	"Api-Go/pkg/storage"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Creacion y manejo de apis con base de datos de postgres")
	db.NewDbPSQL()
	serviceU := user.NewServicec(storage.NewBdUser(db.Pool()))

	if err := serviceU.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}
}
