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

func userAvatar(c *gin.Context) {
	jwtdata, _ := c.Get("data")

	avatar, err := db.GetUserAvatar(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"bg_color": avatar.BgColor,
		"face_color": avatar.FaceColor,
		"face_x": avatar.FaceX,
		"face_y": avatar.FaceY,
		"left_eye_x": avatar.LeX,
		"left_eye_y": avatar.LeY,
		"right_eye_x": avatar.ReX,
		"right_eye_y": avatar.ReY,
		"bezier": avatar.Bezier,
	})
}
