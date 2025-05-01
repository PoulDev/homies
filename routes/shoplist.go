package routes

import (
	"github.com/PoulDev/roommates-api/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)


func myShoppingCart(c *gin.Context) {
	jwtdata, _ := c.Get("data")
	
	cart, err := db.GetCart(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"items": cart.Items,
	});
}

