package handler

import (
	"Api-Go/pkg/model"
)

type Storage interface {
	Create(u *model.User) error
	Update(ID int, u *model.User) error
	Delete(ID int) error
	GetById(ID int) (*model.User, error)
	GetAll() (model.Users, error)
}
