package jwt

import (
	"fmt"
	"testing"
	"time"
)

const secretKey string = "secret"

var TTL time.Duration = time.Duration(time.Hour * 24 * 7)

func createTokenPairWithClaims() (TokenPair, error) {
	userEmail := "test@test.com"

	tokenPair, err := CreateTokenPair(userEmail, secretKey, TTL, TTL)

	fmt.Println("\ncreated token pair: ", tokenPair)
	fmt.Println()

	return tokenPair, err
}

func TestCreateTokenPair(t *testing.T) {
	_, err := createTokenPairWithClaims()

	if err != nil {
		t.Errorf("CreateJWT() error = %v", err)
	}
}

func TestValidateAccessJWT(t *testing.T) {
	tokenPair, err := createTokenPairWithClaims()

	if err != nil {
		t.Errorf("\nCreateJWT() error = %v", err)
	}

	claims, err := VerifyJWT(tokenPair.AccessToken, secretKey)

	if err != nil {
		t.Errorf("ValidateAccessJWT() error = %v", err)
	}

	fmt.Println("\nparsed claims: ", claims)
	fmt.Println()
}

func TestValidateRefreshJWT(t *testing.T) {
	tokenPair, err := createTokenPairWithClaims()

	if err != nil {
		t.Errorf("\nCreateJWT() error = %v", err)
	}

	claims, err := VerifyJWT(tokenPair.RefreshToken, secretKey)

	if err != nil {
		t.Errorf("ValidateRefreshJWT() error = %v", err)
	}

	fmt.Println("\nparsed claims: ", claims)
	fmt.Println()
}

func TestRefreshTokenPair(t *testing.T) {
	tokenPair, err := createTokenPairWithClaims()

	if err != nil {
		t.Errorf("\nCreateJWT() error = %v", err)
	}

	newTokenPair, err := RefreshTokens(
		tokenPair.RefreshToken,
		secretKey,
		TTL,
	)

	if tokenPair.AccessToken == newTokenPair.AccessToken {
		t.Errorf("Access token similar to old")
	}
}
