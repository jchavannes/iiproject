package db

type Session struct {
	Id       uint `gorm:"primary_key"`
	CookieId string `gorm:"unique_index"`
	UserId   uint
	StartTs  uint
}

func (s *Session) Find() error {
	result := find(s, s)
	if result.Error != nil && result.Error.Error() == "record not found" {
		result = create(s)
	}
	return result.Error
}

func (s *Session) Save() error {
	result := save(&s)
	return result.Error
}
