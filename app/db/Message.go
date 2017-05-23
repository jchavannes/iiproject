package db

import (
	"time"
	"fmt"
)

type Message struct {
	Id        uint `gorm:"primary_key"`
	UserId    uint
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

func GetMessages(userId uint) ([]*Message, error) {
	var messages []*Message
	result := find(&messages, &Message{
		UserId: userId,
	})
	if result.Error != nil {
		return nil, fmt.Errorf("Unable to get messages: %s", result.Error)
	}
	return messages, nil
}