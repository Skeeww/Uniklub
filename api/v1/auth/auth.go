package auth

import (
	"context"
	"net/mail"

	"github.com/gin-gonic/gin"
	"noan.dev/uniklub/auth"
	"noan.dev/uniklub/constants"
)

func Login(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		type loginRequestPost struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		var request loginRequestPost
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, constants.CreateErrorMessage(err.Error()))
			return
		}

		address, err := mail.ParseAddress(request.Email)
		if err != nil {
			c.JSON(400, constants.CreateErrorMessage(constants.WrongMailFormat))
			return
		}

		authenticator := ctx.Value(constants.AuthContext).(auth.Authenticater)
		token, err := authenticator.Auth(&auth.UserPasswordCrendentials{
			Email:    address.Address,
			Password: request.Password,
		})
		if err != nil {
			c.JSON(403, constants.CreateErrorMessage(err.Error()))
			return
		}

		c.JSON(200, gin.H{
			"token": token.ToString(),
		})
	}
}

func Me(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		authenticator := ctx.Value(constants.AuthContext).(auth.Authenticater)
		token := c.Keys[constants.TokenContext].(auth.Tokener)

		u, err := authenticator.GetUser(token)
		if err != nil {
			c.JSON(500, constants.CreateErrorMessage(err.Error()))
			return
		}

		c.JSON(200, u)
	}
}
