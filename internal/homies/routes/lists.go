package routes

import (
	"github.com/PoulDev/roommates-api/internal/homies/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type ItemInput struct {
	Text string 	`json:"text"`
}

func getLists(c *gin.Context) {
	jwtdata, _ := c.Get("data")
	
	lists, err := db.GetLists(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"items": lists,
	});
}

func newItem(c *gin.Context) {
	jwtdata, _ := c.Get("data")

	var item ItemInput;
	err := c.ShouldBind(&item);
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON Data!"})
		return
	}

	id_param := c.Param("id")

	err = db.NewItem(item.Text, id_param, jwtdata.(jwt.MapClaims)["uid"].(string));
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{});

}

func getItems(c *gin.Context) {
	id_param := c.Param("id")

	items, err := db.GetItems(id_param)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"items": items})
}

