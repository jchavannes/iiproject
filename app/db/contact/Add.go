package contact

import (
	"github.com/jchavannes/iiproject/app/db"
)

func AddContact(eid string, publicKey string, userId uint) error {
	contact, err := db.GetContact(eid)
	var contactId uint
	if err != nil {
		contactId, err = db.AddContact(eid, publicKey)
		if err != nil {
			return err
		}
	} else {
		contactId = contact.Id
	}
	return db.AddUserContact(userId, contactId)
}
