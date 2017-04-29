package profile

import (
	"github.com/jchavannes/iiproject/db"
)

func Edit(userId uint, profileString string) error {
	user, err := db.GetUserById(userId)
	if err != nil {
		return err
	}
	profile := db.Profile{
		UserId: userId,
		User: user,
		Profile: profileString,
	}
	err = profile.Save()
	return err
}
