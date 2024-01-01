package invoice

import (
	invoiceitem "go-postgres/pkg/InvoiceItem"
	invoiceheader "go-postgres/pkg/invoiceHeader"
)

type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

type Storage interface {
	Create(*Model) error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
