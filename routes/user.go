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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"username": user.Username,
		"email": user.Email,
		"house": db.IdOrnull(user.House),
		"avatar": user.Avatar,
	});
}

func userHouse(c *gin.Context) {
	jwtdata, _ := c.Get("data")

	house, err := db.GetUserHouse(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"name": house.Name,
		"owner": house.Owner,
		"members": house.Members,
	})
}
