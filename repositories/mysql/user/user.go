package user

import (
	userEntities "backend-mental-guardians/entities/user"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (u *UserRepo) FindByEmail(email string) (userEntities.User, error) {
	var userDB User

	if err := u.DB.Where("email = ?", email).First(&userDB).Error; err != nil {
		return userEntities.User{}, err
	}

	return userEntities.User{
		ID:        userDB.ID,
		Email:     userDB.Email,
		Password:  userDB.Password,
		FirstName: userDB.FirstName,
		LastName:  userDB.LastName,
	}, nil
}

func (u *UserRepo) Create(user userEntities.User) (userEntities.User, error) {
	userDB := User{
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	if err := u.DB.Create(&userDB).Error; err != nil {
		return userEntities.User{}, err
	}

	newUser := userEntities.User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return newUser, nil
}