package user

import (
	"errors"
	"fmt"
	"github.com/jchavannes/iiproject/app/db"
	"golang.org/x/crypto/bcrypt"
)

func Signup(cookieId string, username string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user, err := db.CreateUser(username, string(hashedPassword))
	if err != nil {
		return errors.New(fmt.Sprintf("Error signing up: %s\n", err))
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
