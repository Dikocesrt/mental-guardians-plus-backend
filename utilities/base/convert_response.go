package base

import (
	"backend-mental-guardians/constants"
	"net/http"
)

func ConvertResponseCode(err error) int {
	switch err {
		case constants.ErrEmptyFieldRegister:
			return http.StatusBadRequest

		case constants.ErrHashPassword:
			return http.StatusInternalServerError

		case constants.ErrEmailAlreadyRegistered:
			return http.StatusBadRequest

		case constants.ErrCreateUser:
			return http.StatusInternalServerError

		case constants.ErrCreateToken:
			return http.StatusInternalServerError

		case constants.ErrEmptyFieldLogin:
			return http.StatusBadRequest

		case constants.ErrUserNotFound:
			return http.StatusNotFound

		default:
			return http.StatusInternalServerError
	}
}