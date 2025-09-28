package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/zibbadies/homies/internal/homies/models"
)

func AdminMiddleware(c *gin.Context) {
	op, _ := c.Get("op")
	if (!op.(bool)) {
		c.JSON(401, gin.H{"error": models.DBError{
			Message:   "You are not authorized to access this",
			ErrorCode: models.NotAuthorized,
		}})
		c.Abort()
		return
	}

	c.Next()
}
