package constants

import "errors"

var ErrEmptyFieldRegister error = errors.New("Email, Password, First Name, and Last Name cannot be empty")

var ErrHashPassword error = errors.New("Failed to hash password")

var ErrEmailAlreadyRegistered error = errors.New("Email already registered")

var ErrCreateUser error = errors.New("Failed to create user")

var ErrCreateToken error = errors.New("Failed to create token")