package invoiceheader

import "time"

// Modelo de invoiceHeader
type Model struct {
	Id        uint      `json:"id"`
	Client    string    `json:"client"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
