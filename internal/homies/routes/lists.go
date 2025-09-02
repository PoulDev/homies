package routes

import (
	"github.com/PoulDev/homies/internal/homies/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"strconv"
)

type ItemInput struct {
	Text string 	`json:"text"`
}

func getLists(c *gin.Context) {
	jwtdata, _ := c.Get("data")
	
	hid, err := strconv.ParseInt(jwtdata.(jwt.MapClaims)["hid"].(string), 10, 64)

	lists, err := db.GetLists(hid)
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

	list_hid, err := db.GetListHID(id_param)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if jwtdata.(jwt.MapClaims)["hid"].(string) != list_hid {
		c.JSON(400, gin.H{"error": "You can't access other people lists!"})
		return
	}

	err = db.NewItem(item.Text, id_param, jwtdata.(jwt.MapClaims)["uid"].(string));
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{});

}

func getItems(c *gin.Context) {
	jwtdata, _ := c.Get("data")
	id_param := c.Param("id")

	list_hid, err := db.GetListHID(id_param)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if jwtdata.(jwt.MapClaims)["hid"].(string) != list_hid {
		c.JSON(400, gin.H{"error": "You can't access other people lists!"})
		return
	}

	items, err := db.GetItems(id_param)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"items": items})
}

