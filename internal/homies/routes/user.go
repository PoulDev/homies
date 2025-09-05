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
	} else { // let's avoid it for now...
		c.JSON(400, gin.H{"error": "You can't see other people info!"})
		return
	}

	dbuser, err := db.GetUser(uid)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dbuser.Account);
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

	overview := models.Overview{
		User: user.User,
		House: house,
		Items: make([]models.Item, 0),
	}

	c.JSON(200, overview);
}
