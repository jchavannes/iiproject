package db

import (
	"time"
	"fmt"
)

type Message struct {
	Id        uint `gorm:"primary_key"`
	User      *User
	UserId    uint
	Contact   *Contact
	ContactId uint
	Outgoing  bool
	Message   string
	SendTime  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Message) Save() error {
	result := save(m)
	return result.Error
}

func (m *Message) GetFormattedDate() string {
	return m.SendTime.Format(time.RFC3339)
}

func GetMessages(userId uint) ([]*Message, error) {
	var messages []*Message
	result := findOrderBy(&messages, "send_time DESC", &Message{
		UserId: userId,
	})
	if result.Error != nil {
		return nil, fmt.Errorf("Unable to get messages: %s", result.Error)
	}
	return messages, nil
}

func DeleteMessage(messageId uint, userId uint) error {
	message := Message{
		Id: messageId,
		UserId: userId,
	}
	result := remove(&message)
	if result.Error != nil {
		return fmt.Errorf("Error deleting message: %s", result.Error)
	}
	return nil
}
