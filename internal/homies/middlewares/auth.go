package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/PoulDev/homies/pkg/homies/auth"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	data, err := auth.CheckToken(token)
	if (err != nil) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Set("data", data)
	

	op := data["op"]
	if (op == nil || !op.(bool)) {
		c.Set("op", false)
	} else {
		c.Set("op", true)
	}

	c.Next()
}
