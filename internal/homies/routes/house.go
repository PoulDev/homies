package routes

import (
	"github.com/zibbadies/homies/internal/homies/db"
	"github.com/zibbadies/homies/internal/homies/checks"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JHouse struct {
	Name string 		`json:"name" binding:"required,min=4"`
}

func createHouse(c *gin.Context) {
	jwtdata, _ := c.Get("data")

	user, err := db.GetUser(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if user.HouseId != "null" {
		c.JSON(400, gin.H{"error": "You already have a house!"})
		return
	}

	var house JHouse;
	err = c.ShouldBind(&house)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON Data!"})
		return
	}

	err = checks.Check("house_name", house.Name)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	houseid, invite, err := db.NewHouse(house.Name)
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
		"invite": invite,
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
		"invite": house.Invite,
		"members": house.Members,
	})
}

func joinHouse(c *gin.Context) {
	jwtdata, _ := c.Get("data")
	invite := c.Param("invite")

	user, err := db.GetUser(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if user.HouseId != "null" {
		c.JSON(400, gin.H{"error": "You already have a house!"})
		return
	}

	houseid, err := db.HouseIDByInvite(invite)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = db.ChangeHouse(jwtdata.(jwt.MapClaims)["uid"].(string), houseid, false);
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{});
}

func inviteInfo(c *gin.Context) {
	invite := c.Param("invite")

	houseid, err := db.HouseIDByInvite(invite)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	house, err := db.GetHouse(houseid, "")
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, house)
}

func leaveHouse(c *gin.Context) {
	jwtdata, _ := c.Get("data")

	err := db.LeaveHouse(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{});
}
