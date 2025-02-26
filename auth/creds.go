package auth

import (
	"noan.dev/uniklub/models/user"
)

type Credentialer interface {
	Check() (*user.User, error)
}
