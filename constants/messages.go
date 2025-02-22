package constants

import "github.com/gin-gonic/gin"

const (
	WrongFormat         = "request is not well formated"
	WrongMailFormat     = "email is not a valid address"
	WrongCredentials    = "wrong email/password credentials"
	PasswordLengthError = "the password should be at least 8 chars long"
	PasswordHashError   = "the password is too long"
	InternalError       = "something went wrong on our side, please retry later"
)

func CreateErrorMessage(details string) gin.H {
	return gin.H{
		"error": details,
	}
}
