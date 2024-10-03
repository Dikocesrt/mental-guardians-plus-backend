package therapist

import "backend-mental-guardians/entities"

type Therapist struct {
	ID                     uint
	Name                   string
	Age                    int
	Specialist             string
	PhotoURL               string
	PhoneNumber            string
	Gender                 string
	Experience             int
	Fee                    int
	PracticeCity           string
	PracticeLocation       string
	BachelorAlmamater      string
	BachelorGraduationYear int
	ConsultationMode       string
}

type RepostitoryInterface interface {
	GetAll(metadata entities.Metadata, specialist string) ([]Therapist, error)
	// GetByID(id uint) (Therapist, error)
}

type UseCaseInterface interface {
	GetAll(metadata entities.Metadata, specialist string) ([]Therapist, error)
	// GetByID(id uint) (Therapist, error)
}