package constants

import "github.com/gin-gonic/gin"

const (
	WrongFormat = "request is not well formated"
)

func CreateErrorMessage(details string) gin.H {
	return gin.H{
		"error": details,
	}
}
