package auth

import (
	"errors"
	"fmt"
	"github.com/jchavannes/iiproject/db"
	"golang.org/x/crypto/bcrypt"
)

func Signup(cookieId string, username string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := db.User{
		Username: username,
		PasswordHash: string(hashedPassword),
	}
	err = user.Create()
	if err != nil {
		return errors.New(fmt.Sprintf("Error signing up: %s\n", err))
	}
	session := db.Session{
		CookieId: cookieId,
	}
	err = session.Find()
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
