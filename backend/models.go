package main


import (
"time"
"gorm.io/gorm"
)


type User struct {
ID uint `gorm:"primaryKey" json:"id"`
Username string `gorm:"unique;not null" json:"username"`
Password string `json:"password"` // store hashed in prod; plain here for demo
Token string `json:"token"`
CartID *uint `json:"cart_id"`
CreatedAt time.Time `json:"created_at"`
}


type Item struct {
ID uint `gorm:"primaryKey" json:"id"`
Name string `json:"name"`
Status string `json:"status"`
CreatedAt time.Time `json:"created_at"`
}


type Cart struct {
ID uint `gorm:"primaryKey" json:"id"`
UserID uint `json:"user_id"`
Name string `json:"name"`
Status string `json:"status"`
CreatedAt time.Time `json:"created_at"`
Items []CartItem `gorm:"foreignKey:CartID" json:"items"`
}


type CartItem struct {
ID uint `gorm:"primaryKey" json:"id"`
CartID uint `json:"cart_id"`
ItemID uint `json:"item_id"`
}


type Order struct {
ID uint `gorm:"primaryKey" json:"id"`
CartID uint `json:"cart_id"`
UserID uint `json:"user_id"`
CreatedAt time.Time `json:"created_at"`
}


// convenience: run migrations
func AutoMigrate(db *gorm.DB) error {
return db.AutoMigrate(&User{}, &Item{}, &Cart{}, &CartItem{}, &Order{})
}