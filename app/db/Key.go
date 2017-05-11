package db

type Key struct {
	Id         uint `gorm:"primary_key"`
	User       *User
	UserId     uint `gorm:"unique_index"`
	PrivateKey []byte
	PublicKey  []byte
}

func (k *Key) Save() error {
	result := save(&k)
	return result.Error
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
