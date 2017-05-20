package db

import (
	"time"
	"strings"
)

type Contact struct {
	Id        uint `gorm:"primary_key"`
	Eid       string `gorm:"unique_index"`
	PublicKey string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func AddContact(eid string, publicKey string) (uint, error) {
	contact := Contact{
		Eid: eid,
		PublicKey: publicKey,
	}
	result := save(&contact)
	if result.Error != nil {
		return 0, result.Error
	}
	return contact.Id, nil
}

func GetContact(eid string) (*Contact, error) {
	contact := &Contact{
		Eid: eid,
	}
	result := find(contact, contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}

func GetContacts(contactIds []string) ([]*Contact, error) {
	var whereIn []string
	for range contactIds {
		whereIn = append(whereIn, "?")
	}
	where := "id IN (" + strings.Join(whereIn, ", ") + ")"
	var contacts []*Contact
	result := findString(&contacts, where, contactIds...)
	if result.Error != nil {
		return nil, result.Error
	}
	return contacts, nil
}
