package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// request body for adding an item to cart
type AddToCartInput struct {
	ItemID uint `json:"item_id"`
}

// POST /carts  - add item to current user's cart
func AddToCart(c *gin.Context) {
	user, ok := getUserFromToken(c)
	if !ok {
		return // response already sent
	}

	var input AddToCartInput
	if err := c.ShouldBindJSON(&input); err != nil || input.ItemID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item_id"})
		return
	}

	// check item exists
	var item Item
	if err := DB.First(&item, input.ItemID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch item"})
		return
	}

	// find or create an "open" cart for this user
	var cart Cart
	err := DB.Where("user_id = ? AND status = ?", user.ID, "open").First(&cart).Error
	if err == gorm.ErrRecordNotFound {
		// create new cart
		cart = Cart{
			UserID: user.ID,
			Name:   "Cart for " + user.Username,
			Status: "open",
		}
		if err := DB.Create(&cart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create cart"})
			return
		}

		// update user's current cart id
		user.CartID = &cart.ID
		if err := DB.Save(user).Error; err != nil {
			// not critical, but good to try
		}
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cart"})
		return
	}

	// add cart item
	cartItem := CartItem{
		CartID: cart.ID,
		ItemID: item.ID,
	}

	if err := DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add item to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "item added to cart",
		"cart_id": cart.ID,
	})
}

// GET /carts  - get current user's cart + items
func GetCart(c *gin.Context) {
	user, ok := getUserFromToken(c)
	if !ok {
		return
	}

	var cart Cart
	err := DB.Where("user_id = ? AND status = ?", user.ID, "open").Preload("CartItems").First(&cart).Error
	if err == gorm.ErrRecordNotFound {
		// no cart yet
		c.JSON(http.StatusOK, gin.H{
			"cart_id": nil,
			"items":   []CartItem{},
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart_id": cart.ID,
		"items":   cart.CartItems,
	})
}
