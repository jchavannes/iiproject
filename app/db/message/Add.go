package message

import (
	"github.com/jchavannes/iiproject/app/db"
	"fmt"
	"github.com/jchavannes/iiproject/eid/api"
	"github.com/jchavannes/iiproject/app/db/contact"
)

func Add(senderUserId uint, contactEid string, messageSend *api.MessageSend, outgoing bool) error {
	recipientContact, err := contact.Get(contactEid)
	if err != nil {
		return fmt.Errorf("Error getting contact from db: %s", err)
	}

	message := db.Message{
		UserId: senderUserId,
		ContactId: recipientContact.Id,
		Outgoing: outgoing,
		SendTime: messageSend.SendTime,
		Message: messageSend.Message,
	}
	err = message.Save()
	if err != nil {
		return fmt.Errorf("Error saving message: %s", err)
	}

	return nil
}
