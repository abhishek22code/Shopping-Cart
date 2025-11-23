package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	var err error

	// connect to SQLite DB
	DB, err = gorm.Open(sqlite.Open("shopping_cart.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// auto-migrate tables
	if err := DB.AutoMigrate(&User{}, &Item{}, &Cart{}, &CartItem{}, &Order{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	r := gin.Default()

	// ---------- CORS MIDDLEWARE ----------
	// For hosting (Render) and local dev, keep it simple.
	// If you want to restrict later, replace "*" with your frontend URL.
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	// ------------------------------------

	// ---- USER ROUTES ----
	r.POST("/users", CreateUser)
	r.GET("/users", ListUsers)
	r.POST("/users/login", LoginUser)

	// ---- ITEM ROUTES ----
	r.POST("/items", CreateItem)
	r.GET("/items", ListItems)

	// ---- CART ROUTES ----
	r.POST("/carts", AddToCart)
	r.GET("/carts", GetCart)

	// ---- ORDER ROUTES ----
	r.POST("/orders", CreateOrder)
	r.GET("/orders", ListOrders)

	// simple test route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ---------- START SERVER ----------
	// For Render/hosting: PORT is provided via env var.
	// For local dev: default to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
