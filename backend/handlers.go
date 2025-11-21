package main


var in struct{
Name string `json:"name"`
ItemIDs []uint `json:"item_ids"`
}
if err := c.ShouldBindJSON(&in); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}
// if user already has a cart, clear it and replace
var cart Cart
if err := db.Where("user_id = ?", user.ID).First(&cart).Error; err == nil {
// delete existing cart items and cart
db.Where("cart_id = ?", cart.ID).Delete(&CartItem{})
db.Delete(&cart)
}
cart = Cart{UserID: user.ID, Name: in.Name, Status: "active", CreatedAt: time.Now()}
db.Create(&cart)
for _, iid := range in.ItemIDs {
db.Create(&CartItem{CartID: cart.ID, ItemID: iid})
}
user.CartID = &cart.ID
db.Save(&user)
c.JSON(http.StatusCreated, cart)
}
}


func ListCarts(db *gorm.DB) gin.HandlerFunc {
return func(c *gin.Context) {
var carts []Cart
db.Preload("Items").Find(&carts)
c.JSON(http.StatusOK, carts)
}
}


// Orders
// POST /orders { "cart_id": 1 }
func CreateOrder(db *gorm.DB) gin.HandlerFunc {
return func(c *gin.Context) {
var in struct{ CartID uint `json:"cart_id"` }
if err := c.ShouldBindJSON(&in); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}
// find cart
var cart Cart
if err := db.Where("id = ?", in.CartID).First(&cart).Error; err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
return
}
order := Order{CartID: cart.ID, UserID: cart.UserID, CreatedAt: time.Now()}
db.Create(&order)
// optionally mark cart as checked out
db.Model(&cart).Update("status", "ordered")
c.JSON(http.StatusCreated, order)
}
}


func ListOrders(db *gorm.DB) gin.HandlerFunc {
return func(c *gin.Context) {
var orders []Order
db.Find(&orders)
c.JSON(http.StatusOK, orders)
}
}