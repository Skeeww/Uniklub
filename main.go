package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	authent "noan.dev/uniklub/auth"
	"noan.dev/uniklub/constants"
)

func main() {
	authenticator := &authent.JWTAuth{
		Issuer:        "uniklub-v1",
		SigningMethod: jwt.SigningMethodHS512,
		SigningKey:    []byte("hello-world"),
		Expiration:    jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.AuthContext, authenticator)

	gin.SetMode(gin.DebugMode)

	r := SetupRouter(ctx)
	driver := SetupDatabase(ctx)
	defer driver.Close(ctx)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
