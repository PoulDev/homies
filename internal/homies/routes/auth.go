package routes

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/PoulDev/homies/pkg/homies/auth"
	"github.com/PoulDev/homies/pkg/homies/avatar"
	"github.com/PoulDev/homies/internal/homies/checks"
	"github.com/PoulDev/homies/internal/homies/db"
	"github.com/PoulDev/homies/internal/homies/logger"
)

type JUser struct {
	Password string `json:"pwd" binding:"required"` // REQUIRED
	Username string `json:"name"` // solo in register
}


// Centralized between authRegister and authLogin
func getJWT(uid string, house string) (string, error) {
	return auth.GenToken(jwt.MapClaims{
		"uid": uid,
		"hid": house,
		"exp": time.Now().UTC().Add(time.Hour * 24 * 21).Unix(),
	})
}

func authRegister(c *gin.Context) {
	var user JUser;
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON Data!"})
		return
	}

	// Input Validation

	err = checks.CheckUsername(user.Username)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = checks.CheckPassword(user.Password)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// DataBase: Registering the user

	uid, err := db.Register(user.Username, user.Password, avatar.RandAvatar())
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// JWT: Generating the token
	tokenString, err := getJWT(uid, "null")

	if (err != nil) {
		logger.Logger.Error("JWT error", "err", err.Error())
		c.JSON(400, gin.H{"error": "Internal error, please try again later"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString});
}

func authLogin(c *gin.Context) {
	var user JUser;
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON Data!"})
		return
	}

	dbuser, err := db.Login(user.Username, user.Password)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := getJWT(dbuser.UID, dbuser.House.ID)

	if (err != nil) {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
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

