package routes

import (
	"github.com/PoulDev/roommates-api/pkg/db"
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
		"username": user.Username,
		"house": user.House,
		"avatar": user.Avatar,
	}

	if (id_param == "me") {
		response["email"] = user.Email
	}

	c.JSON(200, response);
}

func getAvatar(c *gin.Context) {
	avatar, err := db.GetAvatar(c.Param("id"))
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
