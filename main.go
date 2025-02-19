package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"noan.dev/uniklub/api/v1/clubs"
	"noan.dev/uniklub/constants"
	"noan.dev/uniklub/database"
	"noan.dev/uniklub/middlewares"
)

func main() {
	ctx := context.Background()

	driver, err := database.Init(ctx, database.ConnectionInformation{
		Username: "postgres",
		Password: "postgres",
		Address:  "127.0.0.1",
		Port:     5432,
		Database: "app",
	})
	if err != nil {
		panic(err)
	}
	ctx = context.WithValue(ctx, constants.DatabaseCtx, driver)

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
		}
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}
