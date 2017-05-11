package db

import "time"

type Profile struct {
	Id        uint `gorm:"primary_key"`
	User      *User
	UserId    uint `gorm:"unique_index"`
	Profile   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Profile) Save() error {
	result := save(&p)
	return result.Error
}

func GetProfileByUserId(userId uint) (string, error) {
	profile := &Profile{
		UserId: userId,
	}
	err := find(profile, profile)
	if err.Error != nil {
		return "", err.Error
	}
	return profile.Profile, nil
}
