package main

import (
	"example/go-auth/initializers"
	"example/go-auth/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func main() {
	router := gin.Default()
	// Routers
	routes.SetupRoutes(router)
	router.Run()
}
