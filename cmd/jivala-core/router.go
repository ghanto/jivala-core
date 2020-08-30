package main

import (
	"github.com/Masterminds/squirrel"
	userApp "github.com/ghanto/jivala-core/pkg/user/application/user"
	"github.com/ghanto/jivala-core/pkg/user/infrastructure/user"
	userInra "github.com/ghanto/jivala-core/pkg/user/interfaces/http/user"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func DefineRoutes(router *gin.Engine, db *sqlx.DB) {
	qb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	userRepo := user.NewUserRepository(db, &qb)
	userService := userApp.NewService(userRepo)
	userHandler := userInra.NewUsersHandler(userService)

	apiRouter := router.Group("/user")
	{
		apiRouter.GET("", userHandler.Get)
		apiRouter.POST("", userHandler.Create)
	}

}
