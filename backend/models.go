package main

import "time"

// Users table
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Token     string // will store login token
	CartID    *uint  // current cart id (optional)
	CreatedAt time.Time
}

// Items table
type Item struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Status    string // e.g. "available"
	CreatedAt time.Time
}

// Carts table
type Cart struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Name      string // optional name
	Status    string // e.g. "open", "ordered"
	CreatedAt time.Time

	CartItems []CartItem
}

// CartItems table
type CartItem struct {
	ID     uint `gorm:"primaryKey"`
	CartID uint `gorm:"not null"`
	ItemID uint `gorm:"not null"`
}

// Orders table
type Order struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint `gorm:"not null"`
	UserID    uint `gorm:"not null"`
	CreatedAt time.Time
}
