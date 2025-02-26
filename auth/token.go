package auth

type Tokener interface {
	GetUserId() (string, error)
	IsValid() bool
	ToString() string
}
