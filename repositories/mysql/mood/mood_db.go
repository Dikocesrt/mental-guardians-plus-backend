package mood

import (
	"backend-mental-guardians/repositories/mysql/user"

	"gorm.io/gorm"
)

type Mood struct {
	gorm.Model
	Content string `gorm:"not null;type:text"`
	IsGood  bool   `gorm:"not null"`
	UserID  int    `gorm:"not null;index"`
	User    user.User `gorm:"foreignKey:UserID"`
}