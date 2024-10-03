package therapist

import (
	"backend-mental-guardians/entities"
	therapistEntities "backend-mental-guardians/entities/therapist"
)

type TherapistUseCase struct {
	therapistRepo therapistEntities.RepostitoryInterface
}

func NewTherapistUseCase(therapistRepo therapistEntities.RepostitoryInterface) *TherapistUseCase {
	return &TherapistUseCase{
		therapistRepo: therapistRepo,
	}
}

func (therapistUseCase *TherapistUseCase) GetAll(metadata entities.Metadata, specialist string) ([]therapistEntities.Therapist, error) {
	therapist, err := therapistUseCase.therapistRepo.GetAll(metadata, specialist)
	if err != nil {
		return []therapistEntities.Therapist{}, err
	}
	return therapist, nil
}