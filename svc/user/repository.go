package user

import "time"

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
