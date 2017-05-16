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

func GetProfileStringByUserId(userId uint) (string, error) {
	profile, err := GetProfileByUserId(userId)
	if err != nil {
		return "", err
	}
	return profile.Profile, nil
}

func GetProfileByUserId(userId uint) (*Profile, error) {
	profile := &Profile{
		UserId: userId,
	}
	err := find(profile, profile)
	if err.Error != nil {
		return nil, err.Error
	}
	return profile, nil
}
