package contact

import (
	"github.com/jchavannes/iiproject/app/db"
)

func DeleteUserContact(userId uint, contactId uint) error {
	contact, err := db.GetUserContact(userId, contactId)
	if err != nil {
		return err
	}
	err = db.DeleteUserContact(contact.Id)
	return err
}
