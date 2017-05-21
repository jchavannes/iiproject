package db

import (
	"time"
	"errors"
)

type UserContact struct {
	Id        uint `gorm:"primary_key"`
	UserId    uint `gorm:"unique_index:user_contact"`
	ContactId uint `gorm:"unique_index:user_contact"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func AddUserContact(userId uint, contactId uint) error {
	userContact := UserContact{
		UserId: userId,
		ContactId: contactId,
	}
	result := save(&userContact)
	return result.Error
}

func GetUserContacts(userId uint) ([]*UserContact, error) {
	var userContacts []*UserContact
	result := find(&userContacts, &UserContact{
		UserId: userId,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return userContacts, nil
}

func GetUserContact(userId uint, contactId uint) (*UserContact, error) {
	userContact := &UserContact{
		UserId: userId,
		ContactId: contactId,
	}
	find(userContact, userContact)
	if userContact.Id == 0 {
		return nil, errors.New("Unable to find record.")
	}
	return userContact, nil
}

func DeleteUserContact(id uint) error {
	userContact := &UserContact{
		Id: id,
	}
	result := remove(userContact)
	return result.Error
}
