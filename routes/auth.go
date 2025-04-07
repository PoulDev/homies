package routes

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/PoulDev/roommates-api/auth"
)

type User struct {
	Email string `json:"email" binding:"required,email"` // REQUIRED
	Password string `json:"pwd" binding:"required"` // REQUIRED
	Name string `json:"name"` // solo in register
}

// TODO: Controllare i campi

func authRegister(c *gin.Context) {
	var user User;
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}


	tokenString, err := auth.GenToken(jwt.MapClaims{
		"uid": "12123", // TODO: user id
		"exp": time.Now().Add(time.Hour * 24 * 7 * 3).Unix(),
	})
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println(tokenString, err)

	c.JSON(200, gin.H{"token": tokenString});
}

func authLogin(c *gin.Context) {
	var user User;
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.String(200, "Successfully logged in!");
}
