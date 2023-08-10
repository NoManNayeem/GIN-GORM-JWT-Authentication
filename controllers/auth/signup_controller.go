// controllers/signup_controller.go
package auth

import (
	"net/http"

	"example/go-auth/initializers"
	"example/go-auth/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Your existing Signup function code
	var body struct {
		Username string
		Password string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Request Body!",
		})
		return
	}

	// Hash Passwords
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Hash Password!",
		})
		return
	}

	// Create User
	user := models.User{Username: body.Username, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Create User!",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Created User",
	})
}
