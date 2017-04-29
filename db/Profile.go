package db

type Profile struct {
	User    *User
	UserId  uint `gorm:"primary_key"`
	Profile string
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
