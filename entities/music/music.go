package music

import "backend-mental-guardians/entities"

type Music struct {
	ID           uint
	Title        string
	Singer       string
	MusicURL     string
	ThumbnailURL string
}

type RepositoryInterface interface {
	GetAll(metadata entities.Metadata) ([]Music, error)
}

type UseCaseInterface interface {
	GetAll(metadata entities.Metadata) ([]Music, error)
}