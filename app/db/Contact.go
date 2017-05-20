package db

import (
	"time"
	"strings"
	"fmt"
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

	contactIdsInterface := make([]interface{}, len(contactIds))
	for i, contactId := range contactIds {
		whereIn = append(whereIn, "?")
		contactIdsInterface[i] = contactId
	}
	where := "id IN (" + strings.Join(whereIn, ", ") + ")"
	var contacts []*Contact
	result := findString(&contacts, where, contactIdsInterface...)
	if result.Error != nil {
		fmt.Printf(" contacts: %#v\n where: %#v\n contactIds: %#v\n", contacts, where, contactIds)
		return nil, result.Error
	}
	return contacts, nil
}
