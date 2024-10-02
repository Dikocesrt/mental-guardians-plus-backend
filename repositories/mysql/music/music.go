package music

import "gorm.io/gorm"

type MusicRepo struct {
	DB *gorm.DB
}

func NewMusicRepo(db *gorm.DB) *MusicRepo {
	return &MusicRepo{
		DB: db,
	}
}