package music

import (
	"backend-mental-guardians/constants"
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
		return []musicEntities.Music{}, constants.ErrMusicNotFound
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

func (m *MusicRepo) GetByID(id uint) (musicEntities.Music, error) {
	musicDB := Music{}
	err := m.DB.Where("id = ?", id).First(&musicDB).Error
	if err != nil {
		return musicEntities.Music{}, constants.ErrMusicNotFound
	}
	music := musicEntities.Music{
		ID:           musicDB.ID,
		Title:        musicDB.Title,
		Singer:       musicDB.Singer,
		MusicURL:     musicDB.MusicURL,
		ThumbnailURL: musicDB.ThumbnailURL,
	}
	return music, nil
}