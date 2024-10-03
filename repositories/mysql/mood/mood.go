package mood

import (
	moodEntities "backend-mental-guardians/entities/mood"

	"gorm.io/gorm"
)

type MoodRepo struct {
	DB *gorm.DB
}

func NewMoodRepo(db *gorm.DB) *MoodRepo {
	return &MoodRepo{
		DB: db,
	}
}

func (m *MoodRepo) Create(mood moodEntities.Mood) (moodEntities.Mood, error) {
	moodDB := &Mood{
		Content: mood.Content,
		IsGood:  mood.IsGood,
		UserID:  mood.User.ID,
	}

	if err := m.DB.Create(moodDB).Error; err != nil {
		return moodEntities.Mood{}, err
	}

	newMood := moodEntities.Mood{
		Content: mood.Content,
		IsGood:  mood.IsGood,
		User:    mood.User,
	}

	return newMood, nil
}