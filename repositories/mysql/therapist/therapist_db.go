package therapist

import "gorm.io/gorm"

type Therapist struct {
	gorm.Model
	Name string `gorm:"not null;type:varchar(255)"`
	Age int `gorm:"not null;type:int"`
	Specialist string `gorm:"not null;type:enum('bullying', 'trauma', 'family', 'school', 'love', 'finance')"`
	PhotoURL string `gorm:"not null;type:varchar(255);name:photo_url"`
	PhoneNumber string `gorm:"not null;type:varchar(255);name:phone_number"`
	Gender string `gorm:"not null;type:enum('male', 'female')"`
	Experience int `gorm:"not null;type:int"`
	Fee int `gorm:"not null;type:int"`
	PracticeCity string `gorm:"not null;type:varchar(255);name:practice_city"`
	PracticeLocation string `gorm:"not null;type:varchar(255);name:practice_location"`
	BachelorAlmamater string `gorm:"not null;type:varchar(255);name:bachelor_almamater"`
	BachelorGraduationYear int `gorm:"not null;type:int;name:bachelor_graduation_year"`
	ConsultationMode string `gorm:"not null;type:enum('online', 'offline', 'both')"`
}