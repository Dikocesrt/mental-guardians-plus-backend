package story

import "gorm.io/gorm"

type StoryRepo struct {
	DB *gorm.DB
}

func NewStoryRepo(db *gorm.DB) *StoryRepo {
	return &StoryRepo{
		DB: db,
	}
}