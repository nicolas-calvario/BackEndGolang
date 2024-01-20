package main

import (
	invoiceitem "go-mysql/pkg/InvoiceItem"
	invoiceheader "go-mysql/pkg/invoiceHeader"
	"go-mysql/pkg/product"
	"go-mysql/pkg/storage"
	"log"
)

func main() {
	storage.NewMysqlDB()
	//migracition()
	CreateProduct()
}

func migracition() {
	serviceProduct := product.NewService(storage.NewMySQLProduct(storage.Pool()))
	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	serviceHeader := invoiceheader.NewService(storage.NewMySQLHeader(storage.Pool()))
	if err := serviceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceHeader.migrate, %v", err)
	}

	serviceInvoiceItem := invoiceitem.NewService(storage.NewMySQLItem(storage.Pool()))
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceItem.migrate, %v", err)
	}
}
func CreateProduct() {
	m := &product.Model{
		Name:         "Curso de Go v2",
		Price:        3434.3,
		Observations: "Segundo curso en go",
	}
	serviceProduct := product.NewService(storage.NewMySQLProduct(storage.Pool()))

	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

}
