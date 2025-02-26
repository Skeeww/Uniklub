package auth

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"noan.dev/uniklub/constants"
	"noan.dev/uniklub/models/user"
)

type JWTToken struct {
	*jwt.Token
	SignedKey []byte
}

func (token *JWTToken) GetUserId() (string, error) {
	return token.Claims.GetSubject()
}

func (token *JWTToken) IsValid() bool {
	return token.Valid
}

func (token *JWTToken) ToString() string {
	if tok, err := token.SignedString(token.SignedKey); err == nil {
		return tok
	}
	return ""
}

func (creds *JWTToken) Check() (*user.User, error) {
	if !creds.IsValid() {
		return nil, fmt.Errorf("provided token is not valid")
	}

	address, err := creds.GetUserId()
	if err != nil {
		return nil, fmt.Errorf("provided token is not able to retrieve user id")
	}

	u, err := user.Find(context.Background(), user.UserPrimaryKey{
		Email: address,
	})
	if err != nil {
		return nil, fmt.Errorf(constants.InternalError)
	}

	return u, nil
}
