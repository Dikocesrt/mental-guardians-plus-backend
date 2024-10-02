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

func (u *UserRepo) FindByEmail(email string) error {
	return u.DB.Where("email = ?", email).First(&User{}).Error
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