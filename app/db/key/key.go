package key

import (
	"github.com/jchavannes/iiproject/app/db"
	"github.com/jchavannes/go-pgp/pgp"
)

func Get(userId uint) (*db.Key, error) {
	key, err := db.GetKeyByUserId(userId)
	if err == nil {
		return key, nil
	}
	user, err := db.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	keyPair, err := pgp.GenerateKeyPair(user.Username, "", user.Username + "@example.com")
	if err != nil {
		return nil, err
	}
	dbKey := db.Key{
		UserId: userId,
		PublicKey: []byte(keyPair.PublicKey),
		PrivateKey: []byte(keyPair.PrivateKey),
	}
	err = dbKey.Save()
	if err != nil {
		return nil, err
	}
	return &dbKey, nil
}
