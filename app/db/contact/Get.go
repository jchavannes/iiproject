package contact

import (
	"github.com/jchavannes/iiproject/app/db"
	"strconv"
	"github.com/jchavannes/iiproject/eid/client"
	"fmt"
)

func GetUserContacts(userId uint) ([]*db.Contact, error) {
	userContacts, err := db.GetUserContacts(userId)
	if err != nil {
		return nil, err
	}
	var contactIds []string
	for _, userContact := range userContacts {
		contactIds = append(contactIds, strconv.Itoa(int(userContact.ContactId)))
	}
	contacts, err := db.GetContacts(contactIds)
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func Get(eid string) (*db.Contact, error) {
	contact, err := db.GetContact(eid)
	if err == nil {
		return contact, nil
	}
	idGetResponse, err := client.GetId(eid)
	if err != nil {
		return nil, fmt.Errorf("Error getting id response: %s", err)
	}
	contact, err = db.AddContact(eid, idGetResponse.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("Error adding contact: %s", err)
	}
	return contact, nil
}
