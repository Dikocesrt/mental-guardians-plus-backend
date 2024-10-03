package video

import (
	"backend-mental-guardians/entities"
	videoEntities "backend-mental-guardians/entities/video"

	"gorm.io/gorm"
)

type VideoRepository struct {
	DB *gorm.DB
}

func NewVideoRepo(db *gorm.DB) *VideoRepository {
	return &VideoRepository{
		DB: db,
	}
}

func (v *VideoRepository) GetAll(metadata entities.Metadata, category string) ([]videoEntities.Video, error) {
	var videos []Video
	query := v.DB.Offset((metadata.Page - 1) * metadata.Limit).Limit(metadata.Limit)

	if category != "" {
		query = query.Where("labels = ?", category)
	}
	err := query.Find(&videos).Error

	if err != nil {
		return []videoEntities.Video{}, err
	}

	videoEnts := make([]videoEntities.Video, len(videos))
	for i, video := range videos {
		videoEnts[i] = videoEntities.Video{
			ID:        video.ID,
			VideoID:   video.VideoID,
			Title:     video.Title,
			Author:    video.Author,
			Views:     video.Views,
			Likes:     video.Likes,
			Comments:  video.Comments,
			Labels:    video.Labels,
			Metadata:  video.Metadata,
			Thumbnail: video.Thumbnail,
		}
	}

	return videoEnts, nil
}