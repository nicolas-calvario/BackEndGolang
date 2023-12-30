package main

import (
	"database/sql"
	"errors"
	"fmt"
	invoiceithem "go-postgres/pkg/InvoiceIthem"
	invoiceheader "go-postgres/pkg/invoiceHeader"
	"go-postgres/pkg/product"
	"go-postgres/pkg/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()
	//migrateTables() 	// metodo para creacion de tablas en  la bd definida

	createProduct() // metodo para guaradar prodcutos
	getall()        // metodo que obtiene un producto por id
	delete()
	getall()
	//getByid()
	//update()
	//getByid()

}

func migrateTables() {
	serviceProduct := product.NewService(storage.NewPsqlProduct(storage.Pool()))
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	serviceHeader := invoiceheader.NewService(storage.NewPsqlInvoiceHeader(storage.Pool()))
	if err := serviceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceHeader.migrate, %v", err)
	}

	serviceInvoiceItem := invoiceithem.NewService(storage.NewPsqlInvoiceItem(storage.Pool()))
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceItem.migrate, %v", err)
	}
}

func createProduct() {
	m := &product.Model{
		Name:         "Curso de Go v2",
		Price:        3434.3,
		Observations: "Segundo curso en go",
	}
	serviceProduct := product.NewService(storage.NewPsqlProduct(storage.Pool()))
	if err := serviceProduct.CreateProduct(m); err != nil {
		log.Fatalf("product.Create %v", err)
	}

}

func getall() {
	serviceProduct := product.NewService(storage.NewPsqlProduct(storage.Pool()))
	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}
	fmt.Println(ms)
}

func getByid() {
	serviceProduct := product.NewService(storage.NewPsqlProduct(storage.Pool()))
	m, err := serviceProduct.GetById(2)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No hay producto con ese id")
	case err != nil:
		log.Fatalf("product.GetById: %v", err)

	default:
		fmt.Println(m)
	}
}
func update() {
	m := &product.Model{
		Id:    2,
		Name:  "Curso testing",
		Price: 150,
	}
	serviceProduct := product.NewService(storage.NewPsqlProduct(storage.Pool()))

	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
}
func delete() {
	serviceProduct := product.NewService(storage.NewPsqlProduct(storage.Pool()))
	err := serviceProduct.Delete(1)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}
