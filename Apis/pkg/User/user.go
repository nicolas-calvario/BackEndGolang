package user

import "time"

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:createdat"`
	UpdatedAt time.Time `json:updatedat"`
}

type IUser interface {
	Migrate() error
	//Create(*User) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetById(uint) (*Model, error)
	// Delete(uint) error
}

type Service struct {
	iUser IUser
}

func NewServicec(i IUser) *Service {
	return &Service{i}
}

func (s *Service) Migrate() error {
	return s.iUser.Migrate()
}

// func (s *Service) Create(u *User) error {
// 	return s.iUser.Create(u)
// }
