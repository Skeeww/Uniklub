package middlewares

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"noan.dev/uniklub/auth"
	"noan.dev/uniklub/constants"
)

func HandleJWTAuthToken(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawToken := c.GetHeader("Authorization")
		if len(rawToken) == 0 {
			c.AbortWithStatus(403)
			return
		}

		authenticator := ctx.Value(constants.AuthContext).(*auth.JWTAuth)
		token, err := jwt.ParseWithClaims(rawToken, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("alg is not valid")
			}
			return authenticator.SigningKey, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(403, constants.CreateErrorMessage(err.Error()))
			return
		}

		jwtToken := &auth.JWTToken{
			Token:     token,
			SignedKey: authenticator.SigningKey,
		}
		if !authenticator.IsValid(jwtToken) {
			c.AbortWithStatusJSON(403, constants.CreateErrorMessage(constants.InvalidToken))
			return
		}

		c.Set(constants.TokenContext, jwtToken)
		c.Next()
	}
}
