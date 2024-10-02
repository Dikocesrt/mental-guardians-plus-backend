package user

type User struct {
	ID        uint
	Email     string
	Password  string
	FirstName string
	LastName  string
	Token     string
}

type RepositoryInterface interface {
	FindByEmail(email string) error
	Create(user User) (User, error)
}

type UseCaseInterface interface {
	Register(user User) (User, error)
}