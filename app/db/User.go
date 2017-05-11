package db

import (
	"time"
)

type User struct {
	Id           uint `gorm:"primary_key"`
	Username     string `gorm:"unique_index"`
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func CreateUser(username string, hashedPassword string) (*User, error) {
	user := &User{
		Username: username,
		PasswordHash: string(hashedPassword),
	}
	result := create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetUserByUsername(username string) (*User, error) {
	user := &User{
		Username: username,
	}
	result := find(user, user)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return user, nil
	}
}

func GetUserById(id uint) (*User, error) {
	user := &User{
		Id: id,
	}
	result := find(user, user)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return user, nil
	}
}
