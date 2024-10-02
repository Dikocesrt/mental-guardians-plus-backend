package music

import (
	"backend-mental-guardians/entities"
	musicEntities "backend-mental-guardians/entities/music"

	"gorm.io/gorm"
)

type MusicRepo struct {
	DB *gorm.DB
}

func NewMusicRepo(db *gorm.DB) *MusicRepo {
	return &MusicRepo{
		DB: db,
	}
}

func (m *MusicRepo) GetAll(metadata entities.Metadata) ([]musicEntities.Music, error) {
	musicDB := []Music{}
	err := m.DB.Offset((metadata.Page-1)*metadata.Limit).Limit(metadata.Limit).Find(&musicDB).Error

	if err != nil {
		return []musicEntities.Music{}, err
	}

	musicEnts := make([]musicEntities.Music, len(musicDB))
	for i, music := range musicDB {
		musicEnts[i] = musicEntities.Music{
			ID:           music.ID,
			Title:        music.Title,
			Singer:       music.Singer,
			MusicURL:     music.MusicURL,
			ThumbnailURL: music.ThumbnailURL,
		}
	}

	return musicEnts, nil
}