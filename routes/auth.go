package routes

import (
	_ "log"
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
		"exp": time.Now().UTC().Add(time.Hour * 24 * 21).Unix(),
	})
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

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

func authRenew(c *gin.Context) {
	jwtdata, _ := c.Get("data")

	exp := time.Unix(int64(jwtdata.(jwt.MapClaims)["exp"].(float64)), 0)
	now := time.Now().UTC()

	// l'utente ha 7 giorni per rinnovare il token
	if exp.Sub(now).Hours() > (24 * 7) {
		c.JSON(400, gin.H{"error": "The token will still be valid for some days, please retry later."})
		return
	}
	// delta < 0 e' controllato dall'auth middleware

	tokenString, err := auth.GenToken(jwt.MapClaims{
		"uid": jwtdata.(jwt.MapClaims)["uid"],
		"exp": time.Now().UTC().Add(time.Hour * 24 * 7 * 3).Unix(),
	})
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": tokenString});
}
