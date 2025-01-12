package mood

import (
	"backend-mental-guardians/constants"
	"backend-mental-guardians/entities"
	moodEntities "backend-mental-guardians/entities/mood"
)

type MoodUseCase struct {
	moodRepository moodEntities.RepositoryInterface
}

func NewMoodUseCase(moodRepository moodEntities.RepositoryInterface) *MoodUseCase {
	return &MoodUseCase{
		moodRepository: moodRepository,
	}
}

func (mu *MoodUseCase) Create(mood moodEntities.Mood) (moodEntities.Mood, error) {
	if mood.Content == "" {
		return moodEntities.Mood{}, constants.ErrEmptyMood
	}
	newMood, err := mu.moodRepository.Create(mood)
	if err != nil {
		return moodEntities.Mood{}, err
	}
	return newMood, nil
}

func (mu *MoodUseCase) GetAllByUserID(id uint, metadata entities.Metadata) ([]moodEntities.Mood, error) {
	moods, err := mu.moodRepository.GetAllByUserID(id, metadata)
	if err != nil {
		return []moodEntities.Mood{}, err
	}
	return moods, nil
}

func (mu *MoodUseCase) GetByID(id uint) (moodEntities.Mood, error) {
	mood, err := mu.moodRepository.GetByID(id)
	if err != nil {
		return moodEntities.Mood{}, err
	}
	return mood, nil
}