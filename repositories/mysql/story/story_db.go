package story

import "gorm.io/gorm"

type Story struct {
	gorm.Model
	Title string `gorm:"not null;type:varchar(255)"`
	Author string `gorm:"not null;type:varchar(255)"`
	Content string `gorm:"not null;type:varchar(255)"`
	ThumbnailURL string `gorm:"not null;type:varchar(255);name:thumbnail_url"`
}