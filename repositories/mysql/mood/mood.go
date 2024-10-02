package mood

import "gorm.io/gorm"

type MoodRepo struct {
	DB *gorm.DB
}

func NewMoodRepo(db *gorm.DB) *MoodRepo {
	return &MoodRepo{
		DB: db,
	}
}