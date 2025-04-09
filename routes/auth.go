package routes

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/PoulDev/roommates-api/pkg/auth"
	"github.com/PoulDev/roommates-api/pkg/db"
)

type User struct {
	Email string `json:"email" binding:"required,email"` // REQUIRED
	Password string `json:"pwd" binding:"required"` // REQUIRED
	Username string `json:"name"` // solo in register
}

// TODO: Controllare i campi
// TODO: Database register & login

func authRegister(c *gin.Context) {
	var user User;
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Data!"})
		return
	}

	uid, err := db.Register(user.Email, user.Username, user.Password)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := auth.GenToken(jwt.MapClaims{
		"uid": uid, // TODO: user id
		"op": true,
		"exp": time.Now().UTC().Add(time.Hour * 24 * 21).Unix(),
	})
	if (err != nil) {
		log.Println("JWT ERROR: ", err.Error())
		c.JSON(400, gin.H{"error": "JWT error"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString});
}

func authLogin(c *gin.Context) {
	var user User;
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Data!"})
		return
	}

	dbuser, err := db.Login(user.Email, user.Password)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := auth.GenToken(jwt.MapClaims{
		"uid": dbuser.UID,
		"exp": time.Now().UTC().Add(time.Hour * 24 * 7 * 3).Unix(),
	})

	c.JSON(200, gin.H{
		"name": dbuser.Username, 
		"email": dbuser.Email,
		"token": tokenString,
	});
}

func authRenew(c *gin.Context) {
	jwtdata, _ := c.Get("data")

	// Controllo che l'account non sia stato rimosso dal DB [ vedi DATABASE.md ]
	_, err := db.GetUser(jwtdata.(jwt.MapClaims)["uid"].(string))
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	exp := time.Unix(int64(jwtdata.(jwt.MapClaims)["exp"].(float64)), 0)
	now := time.Now().UTC()

	// l'utente ha 7 giorni per rinnovare il token
	if exp.Sub(now).Hours() > (24 * 7) {
		c.JSON(400, gin.H{"error": "This token will still be valid for some days, please retry later."})
		return
	}
	// delta < 0 e' controllato dall'auth middleware

	tokenString, err := auth.GenToken(jwt.MapClaims{
		"uid": jwtdata.(jwt.MapClaims)["uid"],
		"exp": time.Now().UTC().Add(time.Hour * 24 * 7 * 3).Unix(),
	})
	if (err != nil) {
		log.Println("JWT ERROR: ", err.Error())
		c.JSON(400, gin.H{"error": "JWT error"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString});
}

