package main


import (
"log"
"math/rand"
"time"


"github.com/gin-gonic/gin"
"gorm.io/driver/sqlite"
"gorm.io/gorm"
)


func main() {
rand.Seed(time.Now().UnixNano())
db, err := gorm.Open(sqlite.Open("ecom.db"), &gorm.Config{})
if err != nil {
log.Fatal(err)
}
AutoMigrate(db)


// seed some items if none
var cnt int64
db.Model(&Item{}).Count(&cnt)
if cnt == 0 {
db.Create(&Item{Name: "Apple"})
db.Create(&Item{Name: "Banana"})
db.Create(&Item{Name: "Carrot"})
}


r := gin.Default()


// public
r.POST("/users", CreateUser(db))
r.GET("/users", ListUsers(db))
r.POST("/users/login", LoginUser(db))


r.POST("/items", CreateItem(db))
r.GET("/items", ListItems(db))


// protected cart endpoints
auth := r.Group("/")
auth.Use(AuthRequired(db))
{
auth.POST("/carts", CreateCart(db))
auth.GET("/carts", ListCarts(db))
auth.POST("/orders", CreateOrder(db))
auth.GET("/orders", ListOrders(db))
}


r.Run(":8080")
}