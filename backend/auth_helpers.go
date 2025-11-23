package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// reads X-Auth-Token header, finds the user, or returns 401
func getUserFromToken(c *gin.Context) (*User, bool) {
	token := c.GetHeader("X-Auth-Token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing auth token"})
		return nil, false
	}

	var user User
	err := DB.Where("token = ?", token).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid auth token"})
		return nil, false
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch user"})
		return nil, false
	}

	return &user, true
}
