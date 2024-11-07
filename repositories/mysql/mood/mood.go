package mood

import (
	"backend-mental-guardians/entities"
	moodEntities "backend-mental-guardians/entities/mood"
	"backend-mental-guardians/entities/user"

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

func (m *MoodRepo) GetAllByUserID(userID uint, metadata entities.Metadata) ([]moodEntities.Mood, error) {
	var moods []Mood
	if err := m.DB.Where("user_id = ?", userID).Order("created_at DESC").Offset((metadata.Page - 1) * metadata.Limit).Limit(metadata.Limit).Find(&moods).Error; err != nil {
		return []moodEntities.Mood{}, err
	}
	moodEnts := make([]moodEntities.Mood, len(moods))
	for i, moodEnt := range moods {
		moodEnts[i] = moodEntities.Mood{
			ID:      moodEnt.ID,
			Content: moodEnt.Content,
			IsGood:  moodEnt.IsGood,
			User:    user.User{
				ID: userID,
			},
			CreatedAt: moodEnt.CreatedAt.String(),
		}
	}
	return moodEnts, nil
}

func (m *MoodRepo) GetByID(id uint) (moodEntities.Mood, error) {
	var mood Mood
	if err := m.DB.Where("id = ?", id).First(&mood).Error; err != nil {
		return moodEntities.Mood{}, err
	}
	return moodEntities.Mood{
		ID:      mood.ID,
		Content: mood.Content,
		IsGood:  mood.IsGood,
		User:    user.User{
			ID: mood.UserID,
		},
		CreatedAt: mood.CreatedAt.String(),
	}, nil
}