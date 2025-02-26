package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"noan.dev/uniklub/models/user"
)

// Create an Authenticater object, T check the identity provided, K is the parsed identity provider
type Authenticater interface {
	Auth(Credentialer) (Tokener, error)
	IsValid(Tokener) bool
	GetUser(Tokener) (*user.User, error)
}

type JWTAuth struct {
	Issuer        string
	SigningMethod jwt.SigningMethod
	SigningKey    []byte
	Expiration    *jwt.NumericDate
}

func (auth *JWTAuth) Auth(creds Credentialer) (Tokener, error) {
	user, err := creds.Check()
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(auth.SigningMethod, jwt.RegisteredClaims{
		Issuer:    auth.Issuer,
		Subject:   user.Email,
		ExpiresAt: auth.Expiration,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	return &JWTToken{
		Token:     token,
		SignedKey: auth.SigningKey,
	}, nil
}

func (auth *JWTAuth) IsValid(token Tokener) bool {
	return token.IsValid()
}

func (auth *JWTAuth) GetUser(token Tokener) (*user.User, error) {
	id, err := token.GetUserId()
	if err != nil {
		return nil, err
	}

	return user.Find(context.Background(), user.UserPrimaryKey{
		Email: id,
	})
}
