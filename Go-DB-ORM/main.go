package main

import (
	"fmt"
	"go-db-orm/pkg/models"
	"go-db-orm/pkg/storage"
)

func main() {
	storage.NewPostgresDB()
	//storage.DB().AutoMigrate(&models.Product{}, &models.InvoiceHeader{}, &models.InvoiceItem{})
	//	ReadAll()
	// ReadOne()
	//Update()
	Delete()
}

func Create() {
	p1 := models.Product{
		Name:  "Curso",
		Price: 2312.3,
	}
	OB := "Comentario de cuerso"
	p1.Observations = &OB
	storage.DB().Create(&p1)
}

func ReadAll() {
	productos := make([]models.Product, 0)
	storage.DB().Find(&productos)

	for _, producto := range productos {
		fmt.Printf("%d - %s \n", producto.ID, producto.Name)
	}

}

func ReadOne() {
	myProduct := models.Product{}

	storage.DB().First(&myProduct, 6)
	fmt.Println(myProduct)
}

func Update() {
	myProduct := models.Product{}

	storage.DB().First(&myProduct, 3)
	fmt.Println(myProduct)
	myProduct.ID = 3

	storage.DB().Model(&myProduct).Updates(models.Product{Name: "Curso de Java", Price: 120})
	storage.DB().First(&myProduct, 3)
	fmt.Println(myProduct)
}

func Delete() {
	//borrado permanenete
	myProductDF := models.Product{}
	myProductDF.ID = 1
	storage.DB().Unscoped().Delete(&myProductDF)

	//borrado suave
	myProductD := models.Product{}
	myProductD.ID = 4

	storage.DB().Delete(&myProductD)
}
