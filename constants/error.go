package constants

import "errors"

var ErrEmptyFieldRegister error = errors.New("email, password, first name and last name cannot be empty")

var ErrHashPassword error = errors.New("failed to hash password")

var ErrEmailAlreadyRegistered error = errors.New("email already registered")

var ErrCreateUser error = errors.New("failed to create user")

var ErrCreateToken error = errors.New("failed to create token")

var ErrEmptyFieldLogin error = errors.New("email and password cannot be empty")

var ErrUserNotFound error = errors.New("email or password is incorrect")

var ErrInvalidToken error = errors.New("invalid token")

var ErrInternalServer error = errors.New("internal server error")

var ErrMusicNotFound error = errors.New("music not found")

var ErrContentNotFound error = errors.New("content not found")