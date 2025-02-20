package users

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/gin-gonic/gin"
	"noan.dev/uniklub/constants"
	"noan.dev/uniklub/models/user"
)

func Create(ctx context.Context) func(*gin.Context) {
	return func(c *gin.Context) {
		type userRequestPost struct {
			Email    string `json:"email" binding:"required"`
			Name     string `json:"name" binding:"required"`
			Surname  string `json:"surname" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		var request userRequestPost
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, constants.CreateErrorMessage(constants.WrongFormat))
			return
		}

		email, err := mail.ParseAddress(request.Email)
		if err != nil {
			c.JSON(400, constants.CreateErrorMessage(constants.WrongMailFormat))
			return
		}

		if len(request.Password) < 8 {
			c.JSON(400, constants.CreateErrorMessage(constants.PasswordLengthError))
			return
		}
		password, err := hashPassword(request.Password)
		if err != nil {
			c.JSON(400, constants.CreateErrorMessage(constants.PasswordHashError))
			return
		}

		u, err := user.Create(ctx, user.UserCreationFields{
			Email:    email.Address,
			Name:     request.Name,
			Surname:  request.Surname,
			Password: password,
		})
		if err != nil {
			if err.Error() == "this email is already used, please select another one" {
				c.JSON(400, constants.CreateErrorMessage(err.Error()))
				return
			}
			fmt.Println("warn:", err.Error())
			c.JSON(500, constants.CreateErrorMessage(constants.InternalError))
			return
		}

		c.JSON(200, u)
	}
}
