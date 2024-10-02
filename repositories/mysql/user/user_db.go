package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"unique;not null;type:varchar(255)"`
	Password string `gorm:"not null;type:varchar(255)"`
	FirstName string `gorm:"not null;type:varchar(255)"`
	LastName string `gorm:"type:varchar(255)"`
}