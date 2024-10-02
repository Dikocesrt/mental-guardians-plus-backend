package content

import (
	"backend-mental-guardians/constants"
	"backend-mental-guardians/entities"
	contentEntities "backend-mental-guardians/entities/content"

	"gorm.io/gorm"
)

type ContentRepo struct {
	DB *gorm.DB
}

func NewContentRepo(db *gorm.DB) *ContentRepo {
	return &ContentRepo{
		DB: db,
	}
}

func (s *ContentRepo) GetAll(metadata entities.Metadata, category string, contentType string) ([]contentEntities.Content, error) {
	storyDB := []Content{}
	query := s.DB.Offset((metadata.Page - 1) * metadata.Limit).Limit(metadata.Limit).Where("type = ?", contentType)

	if category != "" {
		query = query.Where("category = ?", category)
	}
	err := query.Find(&storyDB).Error

	if err != nil {
		return []contentEntities.Content{}, constants.ErrContentNotFound
	}

	storyEnts := make([]contentEntities.Content, len(storyDB))
	for i, story := range storyDB {
		storyEnts[i] = contentEntities.Content{
			ID:           story.ID,
			Title:        story.Title,
			Author:       story.Author,
			Content:      story.Content,
			Category:     story.Category,
			Type:         story.Type,
			ThumbnailURL: story.ThumbnailURL,
		}
	}
	return storyEnts, nil
}

func (s *ContentRepo) GetByID(id uint) (contentEntities.Content, error) {
	storyDB := Content{}

	err := s.DB.Where("id = ?", id).First(&storyDB).Error
	if err != nil {
		return contentEntities.Content{}, constants.ErrContentNotFound
	}

	return contentEntities.Content{
		ID:           storyDB.ID,
		Title:        storyDB.Title,
		Author:       storyDB.Author,
		Content:      storyDB.Content,
		Category:     storyDB.Category,
		ThumbnailURL: storyDB.ThumbnailURL,
		Type:         storyDB.Type,
	}, nil
}