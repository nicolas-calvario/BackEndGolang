package user

import "time"

type Model struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:createdat"`
	UpdatedAt time.Time `json:updatedat"`
}

type IUser interface {
	Migrate() error
	Create(*Model) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetById(uint) (*Model, error)
	// Delete(uint) error
}

type Service struct {
	iUser IUser
}

func NewService(i IUser) *Service {
	return &Service{i}
}

func (s *Service) Migrate() error {
	return s.iUser.Migrate()
}

func (s *Service) Create(u *Model) error {
	u.CreatedAt = time.Now()
	return s.iUser.Create(u)
}
