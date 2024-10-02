package content

import "backend-mental-guardians/entities"

type Content struct {
	ID           uint
	Title        string
	Author       string
	Content      string
	Category     string
	Type         string
	ThumbnailURL string
}

type RepositoryInterface interface {
	GetAll(metadata entities.Metadata, category string, contentType string) ([]Content, error)
	GetByID(id uint) (Content, error)
}

type UseCaseInterface interface {
	GetAll(metadata entities.Metadata, category string, contentType string) ([]Content, error)
	GetByID(id uint) (Content, error)
}