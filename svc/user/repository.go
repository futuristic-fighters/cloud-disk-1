package user

import (
	"errors"
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Password   string    `json:"-"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Type       int       `json:"type"`
	Avatar     string    `json:"avatar"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

type Repository interface {
	Find(ID int) *User
}

func (u *User) CheckName() error {
	if u.Name == "" {
		return errors.New("empty name")
	}

	return nil
}

func (u *User) CheckType() error {
	if u.Type != 0 && u.Type != 1 {
		return errors.New("invalid type")
	}

	return nil
}

//...
