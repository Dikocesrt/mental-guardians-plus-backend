package therapist

import (
	"backend-mental-guardians/entities"
	therapistEntities "backend-mental-guardians/entities/therapist"

	"gorm.io/gorm"
)

type TherapistRepo struct {
	DB *gorm.DB
}

func NewTherapistRepo(db *gorm.DB) *TherapistRepo {
	return &TherapistRepo{
		DB: db,
	}
}

func (t *TherapistRepo) GetAll(metadata entities.Metadata, specialist string) ([]therapistEntities.Therapist, error) {
	var therapists []Therapist
	query := t.DB.Offset((metadata.Page - 1) * metadata.Limit).Limit(metadata.Limit)
	
	if specialist != "" {
		query = t.DB.Where("specialist = ?", specialist)
	}
	
	err := query.Find(&therapists).Error

	if err != nil {
		return []therapistEntities.Therapist{}, err
	}

	therapistEnts := make([]therapistEntities.Therapist, len(therapists))
	for i, therapist := range therapists {
		therapistEnts[i] = therapistEntities.Therapist{
			ID:           therapist.ID,
			Name:         therapist.Name,
			Age:          therapist.Age,
			Specialist:  therapist.Specialist,
			PhotoURL:     therapist.PhotoURL,
			PhoneNumber: therapist.PhoneNumber,
			Gender:       therapist.Gender,
			Experience:   therapist.Experience,
			Fee:          therapist.Fee,
			PracticeCity: therapist.PracticeCity,
			PracticeLocation: therapist.PracticeLocation,
			BachelorAlmamater: therapist.BachelorAlmamater,
			BachelorGraduationYear: therapist.BachelorGraduationYear,
			ConsultationMode: therapist.ConsultationMode,
		}
	}

	return therapistEnts, nil
}