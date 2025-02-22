package auth

import (
	"bytes"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Create an Authenticater object, T check the identity provided
type Authenticater[T Credentialer] interface {
	Auth(T) (string, error)
}

type JWTAuth[T Credentialer] struct {
	Issuer        string
	SigningMethod jwt.SigningMethod
	SigningKey    string
	Expiration    *jwt.NumericDate
}

func (auth *JWTAuth[T]) Auth(creds T) (string, error) {
	user, err := creds.Check()
	if err != nil {
		return "", err
	}

	token, err := jwt.NewWithClaims(auth.SigningMethod, jwt.RegisteredClaims{
		Issuer:    auth.Issuer,
		Subject:   user.Email,
		ExpiresAt: auth.Expiration,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}).SignedString(bytes.NewBufferString(auth.SigningKey).Bytes())

	if err != nil {
		return "", err
	}

	return token, nil
}
