package auth

import (
	"errors"
	"fmt"
	"github.com/jchavannes/iiproject/db"
	"golang.org/x/crypto/bcrypt"
)

func Login(cookieId string, username string, password string) error {
	user, err := db.GetUserByUsername(username)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return err
	}

	session, err := db.GetSession(cookieId)
	if err != nil {
		return errors.New(fmt.Sprintf("Error getting session: %s\n", err))
	}

	session.UserId = user.Id
	err = session.Save()
	if err != nil {
		return errors.New(fmt.Sprintf("Error saving session: %s\n", err))
	}

	return nil
}
