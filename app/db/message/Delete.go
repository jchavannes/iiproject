package message

import (
	"github.com/jchavannes/iiproject/app/db"
)

func Delete(messageId uint, userId uint) error {
	return db.DeleteMessage(messageId, userId)
}
