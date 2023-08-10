package main

import (
	"example/go-auth/controllers"
	"example/go-auth/initializers"
	"example/go-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()
	// Sign Up
	r.POST("/signup", controllers.Signup)
	// Login
	r.POST("/login", controllers.Login)
	// Validate
	r.GET("/protected", middlewares.LoginRequired, controllers.Validate)
	r.Run()
}
