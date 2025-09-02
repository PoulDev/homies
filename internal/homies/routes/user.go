package routes

import (
	"github.com/PoulDev/homies/internal/homies/db"
	"github.com/PoulDev/homies/internal/homies/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func userInfo(c *gin.Context) {
	id_param := c.Param("id")
	jwtdata, _ := c.Get("data")

	uid := id_param

	if (id_param == "me") {
		uid = jwtdata.(jwt.MapClaims)["uid"].(string)
	}

	user, err := db.GetUser(uid)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"name": user.Username,
		"house": user.House,
		"avatar": user.Avatar,
	}

	/*
	// For additional information provided if the requested user is also the user making the request

	if (id_param == "me") {

	}
	*/

	c.JSON(200, response);
}

func homeOverview(c *gin.Context) {
	jwtdata, _ := c.Get("data")	

	user, err := db.GetUser(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	uid, err := db.UUIDString2Bytes(user.UID)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	house, err := db.GetHouse(jwtdata.(jwt.MapClaims)["hid"].(string), uid)

	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.House = house

	overview := models.Overview{
		User: user,
		Items: make([]models.Item, 0),
	}

	c.JSON(200, overview);
}
