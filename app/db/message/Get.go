package message

import (
	"github.com/jchavannes/iiproject/app/db"
	"fmt"
)

func GetMessages(userId uint) ([]*db.Message, error) {
	fmt.Printf("UserId: %d\n", userId)
	messages, err := db.GetMessages(userId)
	if err != nil {
		return nil, fmt.Errorf("Error getting messages from db: %s", err)
	}
	users := make(map[uint]*db.User)
	contacts := make(map[uint]*db.Contact)
	for _, message := range messages {
		if _, ok := users[message.UserId]; ! ok {
			users[message.UserId], err = db.GetUserById(message.UserId)
			if err != nil {
				return nil, fmt.Errorf("Error getting user from db: %s", err)
			}
		}
		if _, ok := contacts[message.ContactId]; ! ok {
			contacts[message.ContactId], err = db.GetContactById(message.ContactId)
			if err != nil {
				return nil, fmt.Errorf("Error getting user from db: %s", err)
			}
		}
		message.User = users[message.UserId]
		message.Contact = contacts[message.ContactId]
	}
	return messages, nil
}
