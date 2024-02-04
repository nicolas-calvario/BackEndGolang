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
	Create()

}

func Migrate() {
	serviceU := user.NewService(storage.NewBdUser(db.Pool()))

	if err := serviceU.Migrate(); err != nil {
		log.Fatalf("user.Migrate: %v", err)
	}
}
func Create() {
	u := &user.Model{
		Name: "Nicolas",
	}
	serviceU := user.NewService(storage.NewBdUser(db.Pool()))

	if err := serviceU.Create(u); err != nil {
		log.Fatalf("user.Create: %v", err)
	}
}
