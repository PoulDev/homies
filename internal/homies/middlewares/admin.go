package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AdminMiddleware(c *gin.Context) {
	op, _ := c.Get("op")
	if (!op.(bool)) {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
