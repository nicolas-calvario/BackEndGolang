package main

import (
	invoiceithem "go-postgres/pkg/InvoiceIthem"
	invoiceheader "go-postgres/pkg/invoiceHeader"
	"go-postgres/pkg/product"
	"go-postgres/pkg/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()
	//migrateTables()
	createProduct()

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
