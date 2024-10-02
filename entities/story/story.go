package story

import "backend-mental-guardians/entities"

type Story struct {
	ID           uint
	Title        string
	Author       string
	Content      string
	Category     string
	ThumbnailURL string
}

type RepositoryInterface interface {
	GetAll(metadata entities.Metadata, category string) ([]Story, error)
}

type UseCaseInterface interface {
	GetAll(metadata entities.Metadata, category string) ([]Story, error)
}