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
	FindByEmail(email string) (User, error)
	Create(user User) (User, error)
}

type UseCaseInterface interface {
	Register(user User) (User, error)
	Login(user User) (User, error)
}