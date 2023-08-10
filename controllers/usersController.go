package controllers

import (
	"example/go-auth/initializers"
	"example/go-auth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Sign Up
func Signup(c *gin.Context) {
	// GET the email/pass from requests body
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

// Login
func Login(c *gin.Context) {
	// GET Request Body
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

	// Look Up Requested User
	var user models.User
	initializers.DB.First(&user, "Username = ?", body.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Credentials!",
		})
		return

	}

	// Compare Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Password!",
		})
		return
	}

	// Generate JWT Token
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 25 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed To Generate Token!",
		})
		return

	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

// Validate
func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"messgae": "Inside Validate Routes",
	})
}
