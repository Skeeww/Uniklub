package auth

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/mail"

	"noan.dev/uniklub/constants"
	"noan.dev/uniklub/models/user"
)

type Credentialer interface {
	Check() (*user.User, error)
}

type UserPasswordCrendentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (creds *UserPasswordCrendentials) Check() (*user.User, error) {
	address, err := mail.ParseAddress(creds.Email)
	if err != nil {
		return nil, fmt.Errorf(constants.WrongMailFormat)
	}

	u, err := user.Find(context.Background(), user.UserPrimaryKey{
		Email: address.Address,
	})
	if err != nil {
		return nil, fmt.Errorf(constants.InternalError)
	}
	if u == nil {
		// Mitigate time based attack
		// FIXME: this solution does not works (~60ms with an existing user while ~15ms without)
		u = &user.User{
			Password: rand.Text(),
		}
	}

	if err := VerifyPassword(creds.Password, u.Password); err != nil {
		return nil, fmt.Errorf(constants.WrongCredentials)
	}

	return u, nil
}
