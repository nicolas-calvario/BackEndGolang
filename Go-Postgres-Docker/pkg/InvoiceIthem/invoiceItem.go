package invoiceithem

import "time"

type Model struct {
	Id              uint      `json:"id"`
	InvoiceHeaderId string    `json:"InvoiceHeaderId"`
	ProductId       uint      `json:"productId"`
	CreatedAt       time.Time `json:"createdAt"`
}

type Storage interface {
	Migrate() error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
