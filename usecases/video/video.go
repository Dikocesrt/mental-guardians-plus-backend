package video

import (
	"backend-mental-guardians/entities"
	videoEntities "backend-mental-guardians/entities/video"
)

type VideoUseCase struct {
	videoRepository videoEntities.RepositoryInterface
}

func NewVideoUseCase(videoRepository videoEntities.RepositoryInterface) *VideoUseCase {
	return &VideoUseCase{
		videoRepository: videoRepository,
	}
}

func (vc *VideoUseCase) GetAll(metadata entities.Metadata, category string) ([]videoEntities.Video, error) {
	videos, err := vc.videoRepository.GetAll(metadata, category)
	if err != nil {
		return []videoEntities.Video{}, err
	}
	return videos, nil
}