package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// helper to generate a random token string
func generateToken() (string, error) {
	bytes := make([]byte, 16) // 16 bytes = 32 hex chars
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// request body for signup
type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// POST /users   - signup
func CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if input.Username == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	user := User{
		Username: input.Username,
		Password: input.Password, // NOTE: in real apps, hash this
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not create user (username might be taken)"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		// we do NOT send password or token back
	})
}

// GET /users   - list all users (for testing)
func ListUsers(c *gin.Context) {
	var users []User
	if err := DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}

	// build a safe response without passwords/tokens
	var result []gin.H
	for _, u := range users {
		result = append(result, gin.H{
			"id":        u.ID,
			"username":  u.Username,
			"cart_id":   u.CartID,
			"createdAt": u.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, result)
}

// request body for login
type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// POST /users/login   - login and return token
func LoginUser(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if input.Username == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	var user User
	err := DB.Where("username = ? AND password = ?", input.Username, input.Password).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login"})
		return
	}

	// generate a new token (overwrites previous one = single active login)
	token, err := generateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	user.Token = token

	if err := DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
