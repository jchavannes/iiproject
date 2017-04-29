package profile

import "github.com/jchavannes/iiproject/db"

func Get(userId uint) (string, error) {
	return db.GetProfileByUserId(userId)
}
