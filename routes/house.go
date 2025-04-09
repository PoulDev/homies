package routes

import (
	"github.com/PoulDev/roommates-api/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type House struct {
	Name string 		`json:"name" binding:"required,min=4"`
}

func createHouse(c *gin.Context) {
	var house House;
	err := c.ShouldBind(&house)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Data!"})
		return
	}

	jwtdata, _ := c.Get("data")
	
	houseid, err := db.NewHouse(house.Name, jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"house": houseid,
	});
}

