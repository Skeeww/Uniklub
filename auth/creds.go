package auth

import (
	"context"
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

	user, err := user.Find(context.Background(), user.UserPrimaryKey{
		Email: address.Address,
	})
	if err != nil {
		return nil, fmt.Errorf(constants.InternalError)
	}
	if user == nil {
		return nil, fmt.Errorf(constants.WrongCredentials)
	}

	if err := VerifyPassword(creds.Password, user.Password); err != nil {
		return nil, fmt.Errorf(constants.WrongCredentials)
	}

	return user, nil
}
