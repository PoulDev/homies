package routes

import (
	"github.com/zibbadies/homies/internal/homies/db"
	"github.com/zibbadies/homies/internal/homies/checks"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type ItemInput struct {
	Text string 	`json:"text"`
}

func getLists(c *gin.Context) {
	jwtdata, _ := c.Get("data")
	
	dbuser, err := db.GetUser(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	lists, err := db.GetLists(dbuser.HouseId)
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

	err = checks.Check("list_item_text", item.Text)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	list_id := c.Param("id")

	list_hid, err := db.GetListHID(list_id)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dbuser, err := db.GetUser(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if dbuser.HouseId != list_hid {
		c.JSON(400, gin.H{"error": "You can't access other people lists!"})
		return
	}

	err = db.NewItem(item.Text, list_id, jwtdata.(jwt.MapClaims)["uid"].(string));
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

	dbuser, err := db.GetUser(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if dbuser.HouseId != list_hid {
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

func updateItem(c *gin.Context) {
	jwtdata, _ := c.Get("data")
	list_id := c.Param("id")
	item_id := c.Param("item_id")

	var item ItemInput;
	err := c.ShouldBind(&item);
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON Data!"})
		return
	}

	err = checks.Check("list_item_text", item.Text)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	/*
		Checks:
		- List ID in user's house (list_hid == user_hid)
		- Item ID in list (item_lid == list_id) ( db side? )
	*/

	dbuser, err := db.GetUser(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	list_hid, err := db.GetListHID(list_id)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if dbuser.HouseId != list_hid {
		c.JSON(400, gin.H{"error": "You can't access this list!"})
		return
	}

	err = db.UpdateItem(list_id, item_id, item.Text, jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{});
}
