package routes

import (
	"github.com/PoulDev/roommates-api/internal/homies/db"
	"github.com/PoulDev/roommates-api/internal/homies/checks"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type House struct {
	Name string 		`json:"name" binding:"required,min=4"`
}

func createHouse(c *gin.Context) {
	if (c.Param("id") != "me") {
		c.JSON(400, gin.H{"error": "You can create houses only for yourself!"})
		return
	}

	var house House;
	err := c.ShouldBind(&house)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON Data!"})
		return
	}

	err = checks.CheckHouseName(house.Name)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jwtdata, _ := c.Get("data")
	
	houseid, err := db.NewHouse(house.Name)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = db.ChangeHouse(jwtdata.(jwt.MapClaims)["uid"].(string), houseid, true);
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"house": houseid,
	});
}

func userHouse(c *gin.Context) {
	if (c.Param("id") != "me") {
		c.JSON(400, gin.H{"error": "You can't see other people houses!"})
		return
	}

	jwtdata, _ := c.Get("data")

	house, err := db.GetUserHouse(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"name": house.Name,
		"members": house.Members,
	})
}
