package db

import (
	"time"
)

type UserContact struct {
	Id        uint `gorm:"primary_key"`
	UserId    uint `gorm:"unique_index:user_contact"`
	ContactId uint `gorm:"unique_index:user_contact"`
	CreatedAt time.Time
	UpdatedAt time.Time
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
