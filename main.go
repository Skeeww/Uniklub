package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"noan.dev/uniklub/api/v1/auth"
	"noan.dev/uniklub/api/v1/clubs"
	"noan.dev/uniklub/api/v1/users"
	authent "noan.dev/uniklub/auth"
	"noan.dev/uniklub/constants"
	"noan.dev/uniklub/database"
	"noan.dev/uniklub/middlewares"
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

	if err := database.CreateConnection(ctx, database.ConnectionInformation{
		Username: "postgres",
		Password: "postgres",
		Address:  "127.0.0.1",
		Port:     5432,
		Database: "app",
	}); err != nil {
		panic(err)
	}
	driver := database.GetConnection()

	defer driver.Close(ctx)

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.Use(middlewares.HandleSecurity(ctx))

	v1Router := r.Group("/v1")
	{
		clubsRouter := v1Router.Group("/clubs")
		{
			clubsRouter.GET("/", clubs.Get(ctx))
			clubsRouter.GET("/:id", clubs.GetById(ctx))
			clubsRouter.POST("/", clubs.Add(ctx))
			clubsRouter.POST("/:id", clubs.Update(ctx))
			clubsRouter.PUT("/:id", clubs.Update(ctx))
			clubsRouter.DELETE("/:id", clubs.Remove(ctx))
		}
		usersRouter := v1Router.Group("/users")
		{
			usersRouter.POST("/", users.Create(ctx))
		}
		authRouter := v1Router.Group("/auth")
		{
			authRouter.POST("/", auth.Login(ctx))
			authRouter.GET("/me", middlewares.HandleJWTAuthToken(ctx), auth.Me(ctx))
			authRouter.GET("/logout")
		}
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
