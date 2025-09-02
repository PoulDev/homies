package routes

import (
	"github.com/PoulDev/roommates-api/internal/homies/db"
	"github.com/gin-gonic/gin"
)

func checkDatabase(c *gin.Context) {
    err := db.CheckConnection()

	status := 200;
	if (err != nil) {status = 400}

	c.JSON(status, gin.H{"error": err});
}

