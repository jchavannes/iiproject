package user

import (
	"errors"
	"fmt"
	"github.com/jchavannes/iiproject/app/db"
)

func Logout(cookieId string) error {
	session, err := db.GetSession(cookieId)
	if err != nil {
		return errors.New(fmt.Sprintf("Error getting session: %s\n", err))
	}

	session.HasLoggedOut = true
	err = session.Save()
	if err != nil {
		return errors.New(fmt.Sprintf("Error saving session: %s\n", err))
	}

	return nil
}
