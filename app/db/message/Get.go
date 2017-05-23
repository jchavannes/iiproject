package message

import (
	"github.com/jchavannes/iiproject/app/db"
	"fmt"
)

func GetMessages(userId uint) ([]*db.Message, error) {
	fmt.Printf("UserId: %d\n", userId)
	return db.GetMessages(userId)
}
