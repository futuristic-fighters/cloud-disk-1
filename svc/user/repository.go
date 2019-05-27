package user

import (
	"cloud-disk/cfg"
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

func (u *User) CheckName() *cfg.Err {
	if u.Name == "" {
		return cfg.NewError(cfg.UserNameValidateErr)
	}

	return nil
}

func (u *User) CheckType() *cfg.Err {

	return nil
}

//...
