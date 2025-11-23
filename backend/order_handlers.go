package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// request body for creating order
type CreateOrderInput struct {
	CartID uint `json:"cart_id"`
}

// POST /orders  - checkout
func CreateOrder(c *gin.Context) {
	user, ok := getUserFromToken(c)
	if !ok {
		return
	}

	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil || input.CartID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart_id"})
		return
	}

	// fetch the cart
	var cart Cart
	err := DB.Where("id = ? AND user_id = ? AND status = ?", input.CartID, user.ID, "open").First(&cart).Error
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no open cart found for this user"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cart"})
		return
	}

	// create new order
	order := Order{
		CartID: input.CartID,
		UserID: user.ID,
	}

	if err := DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	// mark cart as ordered
	cart.Status = "ordered"
	if err := DB.Save(&cart).Error; err != nil {
		// not fatal, but good to try
	}

	// clear user's current cart (optional but nice)
	user.CartID = nil
	DB.Save(user)

	c.JSON(http.StatusOK, gin.H{
		"message":  "order created successfully",
		"order_id": order.ID,
	})
}

// GET /orders - list user's orders
func ListOrders(c *gin.Context) {
	user, ok := getUserFromToken(c)
	if !ok {
		return
	}

	var orders []Order
	if err := DB.Where("user_id = ?", user.ID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
