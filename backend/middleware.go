package main


import (
"net/http"
"strings"


"github.com/gin-gonic/gin"
)


// simple token auth: expect header Authorization: Bearer <token>
func AuthRequired(db *gorm.DB) gin.HandlerFunc {
return func(c *gin.Context) {
auth := c.GetHeader("Authorization")
if auth == "" {
c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
return
}
parts := strings.SplitN(auth, " ", 2)
if len(parts) != 2 || parts[0] != "Bearer" {
c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
return
}
token := parts[1]
var u User
if err := db.Where("token = ?", token).First(&u).Error; err != nil {
c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
return
}
// store user in context
c.Set("user", u)
c.Next()
}
}