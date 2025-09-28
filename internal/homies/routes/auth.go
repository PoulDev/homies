package routes

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/zibbadies/homies/internal/homies/checks"
	"github.com/zibbadies/homies/internal/homies/config"
	"github.com/zibbadies/homies/internal/homies/db"
	"github.com/zibbadies/homies/internal/homies/logger"
	"github.com/zibbadies/homies/internal/homies/models"
	"github.com/zibbadies/homies/pkg/homies/auth"
	"github.com/zibbadies/homies/pkg/homies/avatar"
)

type JUser struct {
	Password string `json:"pwd" binding:"required"` // REQUIRED
	Username string `json:"name"` // solo in register
}


// Centralized between authRegister and authLogin
func getJWT(uid string) (string, error) {
	return auth.GenToken(jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().UTC().Add(config.AT_DAYS).Unix(),
	})
}

func authRegister(c *gin.Context) {
	var user JUser;
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": models.DBError{
			Message: "Invalid JSON Data!",
			ErrorCode: models.JsonFormatError,
		}})
		return
	}

	// Input Validation

	err = checks.Check("username", user.Username)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err})
		return
	}

	err = checks.Check("password", user.Password)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err})
		return
	}

	// DataBase: Registering the user

	logger.Logger.Info("Registering user", "username", user.Username)

	uid, err := db.Register(user.Username, user.Password, avatar.RandAvatar())
	if (err != nil) {
		c.JSON(400, gin.H{"error": err})
		return
	}

	// JWT: Generating the token
	tokenString, err := getJWT(uid)

	if (err != nil) {
		logger.Logger.Error("JWT error", "err", err.Error())
		c.JSON(500, models.DBError{
			Message: "Internal error, please try again later",
			ErrorCode: models.InternalError,
		})
		return
	}

	c.JSON(200, gin.H{"token": tokenString});
}

func authLogin(c *gin.Context) {
	var user JUser;
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": models.DBError{
			Message: "Invalid JSON Data!",
			ErrorCode: models.JsonFormatError,
		}})
		return
	}

	dbuser, err := db.Login(user.Username, user.Password)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err})
		return
	}

	tokenString, err := getJWT(dbuser.UID)

	if (err != nil) {
		c.JSON(400, gin.H{"error": gin.H{"message": err.Error()}})
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
		c.JSON(400, gin.H{"error": err})
		return
	}

	exp := time.Unix(int64(jwtdata.(jwt.MapClaims)["exp"].(float64)), 0)
	now := time.Now().UTC()

	// l'utente ha 7 giorni per rinnovare il token
	if exp.Sub(now).Hours() > (24 * 7) {
		// This error shouldn't be formatted like this
		// but this endpoint will be fucked up anyway sooner or later.
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

