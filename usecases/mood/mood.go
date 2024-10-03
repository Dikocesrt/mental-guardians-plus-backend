package mood

import moodEntities "backend-mental-guardians/entities/mood"

type MoodUseCase struct {
	moodRepository moodEntities.RepositoryInterface
}

func NewMoodUseCase(moodRepository moodEntities.RepositoryInterface) *MoodUseCase {
	return &MoodUseCase{
		moodRepository: moodRepository,
	}
}

func (mu *MoodUseCase) Create(mood moodEntities.Mood) (moodEntities.Mood, error) {
	newMood, err := mu.moodRepository.Create(mood)
	if err != nil {
		return moodEntities.Mood{}, err
	}
	return newMood, nil
}

func (mu *MoodUseCase) GetAllByUserID(id uint) ([]moodEntities.Mood, error) {
	moods, err := mu.moodRepository.GetAllByUserID(id)
	if err != nil {
		return []moodEntities.Mood{}, err
	}
	return moods, nil
}