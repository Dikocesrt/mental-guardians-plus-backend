package music

import "gorm.io/gorm"

type Music struct {
	gorm.Model
	Title string `gorm:"not null;type:varchar(255)"`
	Singer string `gorm:"not null;type:varchar(255)"`
	MusicURL string `gorm:"not null;type:varchar(255);name:music_url"`
	ThumbnailURL string `gorm:"not null;type:varchar(255);name:thumbnail_url"`
}