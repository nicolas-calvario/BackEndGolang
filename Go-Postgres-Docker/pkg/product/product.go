package product

import "time"

// Modelo de producto
type Model struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name"`
	Observations string    `json:observations"`
	Price        float64   `json:price"`
	CreatedAt    time.Time `json:createdat"`
	UpdatedAt    time.Time `json:updatedat"`
}

type Models []*Model
type Storage interface {
	Migrate() error
	Create(*Model) error
	// 	Update(*Model) error
	// 	GetAll() (Models, error)
	// 	GetById(uint) (*Model, error)
	// 	Delete(uint) error
	//
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
func (s *Service) CreateProduct(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}
