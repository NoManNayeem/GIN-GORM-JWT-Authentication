// routes.go
package routes

import (
	"example/go-auth/controllers/auth"
	"example/go-auth/controllers/protected"
	"example/go-auth/controllers/public"
	"example/go-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", auth.Signup)
		authGroup.POST("/login", auth.Login)
	}

	publicGroup := router.Group("/public")
	{
		publicGroup.GET("/", public.PublicRoute)
	}

	protectedGroup := router.Group("/protected")
	{
		protectedGroup.Use(middlewares.LoginRequired)
		protectedGroup.GET("/users", protected.UserList)
		protectedGroup.PATCH("/edit-user/:id", protected.EditUser)
	}
}
