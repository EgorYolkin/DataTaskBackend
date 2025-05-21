package users_handler

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrIncorrectAuthData = errors.New("incorrect auth data")
)
