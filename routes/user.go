package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)


func userMe(c *gin.Context) {
	jwtdata, _ := c.Get("data")
	c.JSON(200, gin.H{"id": jwtdata.(jwt.MapClaims)["uid"]});
}

