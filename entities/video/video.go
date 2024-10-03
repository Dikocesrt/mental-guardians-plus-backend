package video

import "backend-mental-guardians/entities"

type Video struct {
	ID        string
	VideoID   string
	Title     string
	Author    string
	Views     int
	Likes     int
	Comments  int
	Labels    string
	Metadata  string
	Thumbnail string
}

type RepositoryInterface interface {
	GetAll(metadata entities.Metadata, category string) ([]Video, error)
}

type UseCaseInterface interface {
	GetAll(metadata entities.Metadata, category string) ([]Video, error)
}