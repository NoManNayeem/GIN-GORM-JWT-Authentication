package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublicRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to public page!"})
}
