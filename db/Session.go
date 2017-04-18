package db

type Session struct {
	Id           uint `gorm:"primary_key"`
	CookieId     string `gorm:"unique_index"`
	HasLoggedOut bool
	UserId       uint
	StartTs      uint
}

func (s *Session) Save() error {
	result := save(&s)
	return result.Error
}

func GetSession(cookieId string) (*Session, error) {
	session := &Session{
		CookieId: cookieId,
	}
	result := find(session, session)
	if result.Error != nil && result.Error.Error() == "record not found" {
		result = create(session)
	}
	if result.Error != nil {
		return nil, result.Error
	} else {
		return session, nil
	}
}
