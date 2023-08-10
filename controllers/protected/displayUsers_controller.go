// controllers/validate_controller.go
package protected

import (
	"example/go-auth/initializers"
	"example/go-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func UserList(c *gin.Context) {
	// Get all records
	var users []models.User
	initializers.DB.Find(&users)

	// Create a slice to hold the response data
	var userResponses []UserResponse

	// Convert fetched user data to the response format
	for _, user := range users {
		userResponse := UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"), // Format the time
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"), // Format the time
		}
		userResponses = append(userResponses, userResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": userResponses,
	})
}
