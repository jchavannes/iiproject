package auth

import "github.com/jchavannes/iiproject/db"

func IsLoggedIn(cookieId string) bool {
	session, err := db.GetSession(cookieId)
	if err != nil {
		return false
	}
	if session.UserId > 0 && ! session.HasLoggedOut {
		return true
	}
	return false
}

func GetSessionUser(cookieId string) *db.User {
	session, err := db.GetSession(cookieId)
	if err != nil || session.UserId == 0 || session.HasLoggedOut {
		return nil
	}
	user, err := db.GetUserById(session.UserId)
	if err != nil {
		return nil
	}
	return user
}
