package music

import (
	"backend-mental-guardians/entities"
	musicEntities "backend-mental-guardians/entities/music"
)

type MusicUseCase struct {
	musicRepo musicEntities.RepositoryInterface
}

func NewMusicUseCase(musicRepo musicEntities.RepositoryInterface) *MusicUseCase {
	return &MusicUseCase{
		musicRepo: musicRepo,
	}
}

func (musicUseCase *MusicUseCase) GetAll(metadata entities.Metadata) ([]musicEntities.Music, error) {
	musics, err := musicUseCase.musicRepo.GetAll(metadata)
	if err != nil {
		return []musicEntities.Music{}, err
	}
	return musics, nil
}