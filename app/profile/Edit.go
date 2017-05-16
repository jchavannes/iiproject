package profile

import (
	"github.com/jchavannes/iiproject/app/db"
)

func Edit(userId uint, profileString string) error {
	profile, err := db.GetProfileByUserId(userId)
	if err != nil {
		profile = &db.Profile{
			UserId: userId,
		}
	}
	profile.Profile = profileString
	err = profile.Save()
	return err
}
