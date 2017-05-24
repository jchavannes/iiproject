package contact

import (
	"github.com/jchavannes/iiproject/app/db"
	"github.com/jchavannes/iiproject/eid"
)

func AddContact(eidString string, publicKey string, userId uint) error {
	eidString = eid.ConvertFullEidUrlIntoShort(eidString)
	var contact *db.Contact
	contact, err := db.GetContact(eidString)
	if err != nil {
		contact, err = db.AddContact(eidString, publicKey)
		if err != nil {
			return err
		}
	}
	return db.AddUserContact(userId, contact.Id)
}
