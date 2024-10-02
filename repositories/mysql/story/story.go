package story

import (
	"backend-mental-guardians/constants"
	"backend-mental-guardians/entities"
	storyEntities "backend-mental-guardians/entities/story"

	"gorm.io/gorm"
)

type StoryRepo struct {
	DB *gorm.DB
}

func NewStoryRepo(db *gorm.DB) *StoryRepo {
	return &StoryRepo{
		DB: db,
	}
}

func (s *StoryRepo) GetAll(metadata entities.Metadata, category string) ([]storyEntities.Story, error) {
	storyDB := []Story{}
	query := s.DB.Offset((metadata.Page - 1) * metadata.Limit).Limit(metadata.Limit)

	if category != "" {
		query = query.Where("category = ?", category)
	}
	err := query.Find(&storyDB).Error

	if err != nil {
		return []storyEntities.Story{}, constants.ErrStoryNotFound
	}

	storyEnts := make([]storyEntities.Story, len(storyDB))
	for i, story := range storyDB {
		storyEnts[i] = storyEntities.Story{
			ID:           story.ID,
			Title:        story.Title,
			Author:       story.Author,
			Content:      story.Content,
			Category:     story.Category,
			ThumbnailURL: story.ThumbnailURL,
		}
	}
	return storyEnts, nil
}

func (s *StoryRepo) GetByID(id uint) (storyEntities.Story, error) {
	storyDB := Story{}

	err := s.DB.Where("id = ?", id).First(&storyDB).Error
	if err != nil {
		return storyEntities.Story{}, constants.ErrStoryNotFound
	}

	return storyEntities.Story{
		ID:           storyDB.ID,
		Title:        storyDB.Title,
		Author:       storyDB.Author,
		Content:      storyDB.Content,
		Category:     storyDB.Category,
		ThumbnailURL: storyDB.ThumbnailURL,
	}, nil
}