package content

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	Title string `gorm:"not null;type:varchar(255)"`
	Author string `gorm:"not null;type:varchar(255)"`
	Content string `gorm:"not null;type:text"`
	Category string `gorm:"not null;type:enum('bullying', 'trauma', 'family', 'school', 'love', 'finance')"`
	Type string `gorm:"not null;type:enum('story', 'article')"`
	ThumbnailURL string `gorm:"not null;type:varchar(255);name:thumbnail_url"`
}