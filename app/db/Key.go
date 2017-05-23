package db

import (
	"time"
	"github.com/jchavannes/iiproject/eid"
)

type Key struct {
	Id         uint `gorm:"primary_key"`
	User       *User
	UserId     uint `gorm:"unique_index"`
	PrivateKey []byte
	PublicKey  []byte
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (k *Key) Save() error {
	result := save(&k)
	return result.Error
}

func (k *Key) GetKeyPair() eid.KeyPair {
	keyPair := eid.KeyPair{
		PublicKey: k.PublicKey,
		PrivateKey: k.PrivateKey,
	}
	return keyPair
}

func GetKeyByUserId(userId uint) (*Key, error) {
	key := &Key{
		UserId: userId,
	}
	err := find(key, key)
	if err.Error != nil {
		return nil, err.Error
	}
	return key, nil
}
