package constants

import "errors"

var ErrEmptyFieldRegister error = errors.New("nama depan, nama belakang, email atau password tidak boleh kosong")

var ErrEmptyMood error = errors.New("deskripsi mood tidak boleh kosong")

var ErrHashPassword error = errors.New("hash password error")

var ErrEmailAlreadyRegistered error = errors.New("email sudah terpakai")

var ErrCreateUser error = errors.New("register gagal")

var ErrCreateToken error = errors.New("gagal membuat token")

var ErrEmptyFieldLogin error = errors.New("email atau password tidak boleh kosong")

var ErrUserNotFound error = errors.New("email atau password salah")

var ErrInvalidToken error = errors.New("token tidak valid")

var ErrInternalServer error = errors.New("internal server error")

var ErrMusicNotFound error = errors.New("music tidak ditemukan")

var ErrContentNotFound error = errors.New("content tidak ditemukan")

var ErrTherapistNotFound error = errors.New("psikolog tidak ditemukan")