package jwt

import (
	"errors"
	jwt "github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type Claims struct {
	UserID    int    `json:"user_id"`
	UserEmail string `json:"user_email"`
	jwt.RegisteredClaims
}

var ValidMethod = jwt.SigningMethodHS256

func makeClaims(userID int, userEmail, subject string, TTL time.Duration) Claims {
	return Claims{
		UserID:    userID,
		UserEmail: userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   subject,
		},
	}
}

func CreateTokenPair(
	userID int,
	userEmail,
	jwtSecretKey string,
	accessTTL,
	refreshTTL time.Duration,
) (TokenPair, error) {
	accessClaims := makeClaims(userID, userEmail, "access", accessTTL)
	accessToken, err := createJWT(jwtSecretKey, accessClaims)
	if err != nil {
		return TokenPair{}, err
	}

	refreshClaims := makeClaims(userID, userEmail, "refresh", refreshTTL)
	refreshToken, err := createJWT(jwtSecretKey, refreshClaims)
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    accessClaims.ExpiresAt.Unix(),
	}, nil
}

func RefreshTokens(refreshToken, jwtSecretKey string, accessTTL time.Duration) (TokenPair, error) {
	claims, err := VerifyJWT(refreshToken, jwtSecretKey)
	if err != nil {
		return TokenPair{}, ErrInvalidToken
	}

	if mapClaims, ok := claims.(jwt.MapClaims); ok {
		userEmail, ok := mapClaims["user_email"].(string)
		userID, ok := mapClaims["user_id"].(int)
		if !ok {
			return TokenPair{}, ErrInvalidToken
		}

		subject, ok := mapClaims["sub"].(string)
		if !ok || subject != "refresh" {
			return TokenPair{}, ErrInvalidRefreshToken
		}

		return CreateTokenPair(userID, userEmail, jwtSecretKey, accessTTL, time.Hour*24*7)
	}

	return TokenPair{}, ErrInvalidRefreshToken
}

func createJWT(jwtSecretKey string, claims Claims) (string, error) {
	token := jwt.NewWithClaims(ValidMethod, claims)
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", ErrJWTSigningError
	}
	return t, nil
}

func VerifyJWT(token, jwtSecretKey string) (jwt.Claims, error) {
	parsedToken, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidToken
			}
			return []byte(jwtSecretKey), nil
		},
		jwt.WithLeeway(time.Second*60),
		jwt.WithValidMethods([]string{ValidMethod.Name}),
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}

	if !parsedToken.Valid {
		return nil, ErrInvalidToken
	}

	if parsedToken.Method.Alg() != ValidMethod.Alg() {
		return nil, ErrInvalidToken
	}

	return parsedToken.Claims, nil
}
