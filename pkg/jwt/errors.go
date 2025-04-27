package jwt

import "errors"

var (
	ErrJWTSigningError     = errors.New("jwt signing error")
	ErrInvalidToken        = errors.New("invalid jwt")
	ErrTokenExpired        = errors.New("jwt token expired")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)
