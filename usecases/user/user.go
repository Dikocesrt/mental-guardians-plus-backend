package user

import (
	"backend-mental-guardians/constants"
	userEntities "backend-mental-guardians/entities/user"
	"backend-mental-guardians/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo userEntities.RepositoryInterface
}

func NewUserUseCase(userRepo userEntities.RepositoryInterface) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (userUseCase *UserUseCase) Register(user userEntities.User) (userEntities.User, error) {
	if user.Email == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		return userEntities.User{}, constants.ErrEmptyFieldRegister
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return userEntities.User{}, constants.ErrHashPassword
	}

	user.Password = string(hashedPassword)

	_, err = userUseCase.userRepo.FindByEmail(user.Email)
	if err == nil {
		return userEntities.User{}, constants.ErrEmailAlreadyRegistered
	}

	newUser, err := userUseCase.userRepo.Create(user)
	if err != nil {
		return userEntities.User{}, constants.ErrCreateUser
	}

	token, err := middlewares.CreateToken(newUser.ID)
	if err != nil {
		return userEntities.User{}, constants.ErrCreateToken
	}

	newUser.Token = token
	return newUser, nil
}

func (userUseCase *UserUseCase) Login(user userEntities.User) (userEntities.User, error) {
	if user.Email == "" || user.Password == "" {
		return userEntities.User{}, constants.ErrEmptyFieldLogin
	}

	userExist, err := userUseCase.userRepo.FindByEmail(user.Email)
	if err != nil {
		return userEntities.User{}, constants.ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(user.Password))
	if err != nil {
		return userEntities.User{}, constants.ErrUserNotFound
	}

	token, err := middlewares.CreateToken(userExist.ID)
	if err != nil {
		return userEntities.User{}, constants.ErrCreateToken
	}

	userExist.Token = token
	return userExist, nil
}
