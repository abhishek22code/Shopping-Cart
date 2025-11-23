package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// request body for creating an item
type CreateItemInput struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

// POST /items  - create item
func CreateItem(c *gin.Context) {
	var input CreateItemInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	// default status if empty
	if input.Status == "" {
		input.Status = "available"
	}

	item := Item{
		Name:   input.Name,
		Status: input.Status,
	}

	if err := DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create item"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        item.ID,
		"name":      item.Name,
		"status":    item.Status,
		"createdAt": item.CreatedAt,
	})
}

// GET /items  - list all items
func ListItems(c *gin.Context) {
	var items []Item

	if err := DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch items"})
		return
	}

	c.JSON(http.StatusOK, items)
}
