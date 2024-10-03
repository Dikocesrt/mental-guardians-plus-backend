package therapist

import (
	"backend-mental-guardians/constants"
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


func (t *TherapistRepo) GetByID(id uint) (therapistEntities.Therapist, error) {
	var therapistDB Therapist
	err := t.DB.Where("id = ?", id).First(&therapistDB).Error
	if err != nil {
		return therapistEntities.Therapist{}, constants.ErrTherapistNotFound
	}
	therapist := therapistEntities.Therapist{
		ID:           therapistDB.ID,
		Name:         therapistDB.Name,
		Age:          therapistDB.Age,
		Specialist:  therapistDB.Specialist,
		PhotoURL:     therapistDB.PhotoURL,
		PhoneNumber: therapistDB.PhoneNumber,
		Gender:       therapistDB.Gender,
		Experience:   therapistDB.Experience,
		Fee:          therapistDB.Fee,
		PracticeCity: therapistDB.PracticeCity,
		PracticeLocation: therapistDB.PracticeLocation,
		BachelorAlmamater: therapistDB.BachelorAlmamater,
		BachelorGraduationYear: therapistDB.BachelorGraduationYear,
		ConsultationMode: therapistDB.ConsultationMode,
	}
	return therapist, nil
}