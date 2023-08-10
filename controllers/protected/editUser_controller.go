// controllers/protected/edit_user_controller.go
package protected

import (
	"example/go-auth/initializers"
	"example/go-auth/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditUserRequest struct {
	Username string `json:"username"`
}

func EditUser(c *gin.Context) {
	// Get user ID from the route parameters
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Fetch the existing user record
	var user models.User
	if result := initializers.DB.First(&user, userID); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind the request body to the EditUserRequest struct
	var editUserRequest EditUserRequest
	if err := c.ShouldBindJSON(&editUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Update the user record with the new username
	user.Username = editUserRequest.Username
	if result := initializers.DB.Save(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
