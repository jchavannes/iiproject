package db

import (
	"errors"
)

type User struct {
	Id           uint `gorm:"primary_key"`
	Username     string `gorm:"unique_index"`
	PasswordHash string
}

func (u *User) Read() error {
	if u.Id == 0 && u.Username == "" {
		return errors.New("Must set either Id or Username.")
	}
	result := find(u, u)
	return result.Error
}

func (u *User) Create() error {
	result := create(u)
	return result.Error
}
