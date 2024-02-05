package model

import "time"

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:createdat"`
	UpdatedAt time.Time `json:updatedat"`
}

type Users []User
