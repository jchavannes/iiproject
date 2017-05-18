package profile

import "github.com/jchavannes/iiproject/app/db"

func Get(userId uint) (string, error) {
	return db.GetProfileStringByUserId(userId)
}
