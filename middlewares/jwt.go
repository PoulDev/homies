package middlewares;

import (
	"github.com/gin-gonic/gin"
	"github.com/PoulDev/roommates-api/auth"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        data, err := auth.CheckToken(token)
        if (err != nil) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Set("data", data)
        c.Next()
    }
}
