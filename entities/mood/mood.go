package mood

import "backend-mental-guardians/entities/user"

type Mood struct {
	ID      uint
	Content string
	IsGood  bool
	User    user.User
}

type RepositoryInterface interface {
	Create(mood Mood) (Mood, error)
	GetAllByUserID(id uint) ([]Mood, error)
}

type UseCaseInterface interface {
	Create(mood Mood) (Mood, error)
	GetAllByUserID(id uint) ([]Mood, error)
}