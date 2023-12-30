package product

import (
	"fmt"
	"strings"
	"time"
)

// Modelo de producto
type Model struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name"`
	Observations string    `json:observations"`
	Price        float64   `json:price"`
	CreatedAt    time.Time `json:createdat"`
	UpdatedAt    time.Time `json:updatedat"`
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5f | %10s | %10s",
		m.Id, m.Name, m.Observations, m.Price,
		m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

// Models slice of Model
type Models []*Model

func (m Models) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%02s | %-20s | %-20s | %5s | %10s | %10s\n",
		"id", "name", "observations", "price", "created_at", "updated_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}

type Storage interface {
	Migrate() error
	Create(*Model) error
	// 	Update(*Model) error
	GetAll() (Models, error)
	GetById(uint) (*Model, error)
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
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

func (s *Service) GetById(id uint) (*Model, error) {
	return s.storage.GetById(id)
}
