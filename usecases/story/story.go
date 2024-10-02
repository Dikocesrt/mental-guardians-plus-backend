package story

import (
	"backend-mental-guardians/entities"
	storyEntities "backend-mental-guardians/entities/story"
)

type StoryUseCase struct {
	storyRepo storyEntities.RepositoryInterface
}

func NewStoryUseCase(storyRepo storyEntities.RepositoryInterface) *StoryUseCase {
	return &StoryUseCase{
		storyRepo: storyRepo,
	}
}

func (s *StoryUseCase) GetAll(metadata entities.Metadata, category string) ([]storyEntities.Story, error) {
	stories, err := s.storyRepo.GetAll(metadata, category)
	if err != nil {
		return []storyEntities.Story{}, err
	}
	return stories, nil
}

func (s *StoryUseCase) GetByID(id uint) (storyEntities.Story, error) {
	story, err := s.storyRepo.GetByID(id)
	if err != nil {
		return storyEntities.Story{}, err
	}
	return story, nil
}