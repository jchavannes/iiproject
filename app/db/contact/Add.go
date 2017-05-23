package contact

import (
	"github.com/jchavannes/iiproject/app/db"
)

func AddContact(eid string, publicKey string, userId uint) error {
	var contact *db.Contact
	contact, err := db.GetContact(eid)
	if err != nil {
		contact, err = db.AddContact(eid, publicKey)
		if err != nil {
			return err
		}
	}
	return db.AddUserContact(userId, contact.Id)
}
