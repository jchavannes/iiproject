package contact

import (
	"github.com/jchavannes/iiproject/app/db"
	"strconv"
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
