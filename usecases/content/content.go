package story

import (
	"backend-mental-guardians/entities"
	contentEntities "backend-mental-guardians/entities/content"
)

type ContentUseCase struct {
	contentRepo contentEntities.RepositoryInterface
}

func NewContentUseCase(contentRepo contentEntities.RepositoryInterface) *ContentUseCase {
	return &ContentUseCase{
		contentRepo: contentRepo,
	}
}

func (s *ContentUseCase) GetAll(metadata entities.Metadata, category string, contentType string) ([]contentEntities.Content, error) {
	stories, err := s.contentRepo.GetAll(metadata, category, contentType)
	if err != nil {
		return []contentEntities.Content{}, err
	}
	return stories, nil
}

func (s *ContentUseCase) GetByID(id uint) (contentEntities.Content, error) {
	story, err := s.contentRepo.GetByID(id)
	if err != nil {
		return contentEntities.Content{}, err
	}
	return story, nil
}