package routes

import (
	"github.com/PoulDev/roommates-api/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)


func userMe(c *gin.Context) {
	jwtdata, _ := c.Get("data")
	
	user, err := db.GetUser(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": db.PreattyError(err)})
		return
	}

	c.JSON(200, gin.H{
		"username": user.Username,
		"email": user.Email,
		"house": db.IdOrnull(user.House),
	});
}

