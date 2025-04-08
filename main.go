package main
//dai
import (
	"log"
	"time"

	"github.com/PoulDev/roommates-api/config"
	"github.com/PoulDev/roommates-api/pkg/db"
	"github.com/PoulDev/roommates-api/routes"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	err := config.LoadConfig();
	if (err != nil) {
		log.Println("!! FROCIAZZO MANCANO LE ENVIRONMENT VARIABLES");
		log.Panic(err);
	}

	err = db.ConnectDatabase()
	if (err != nil) {
		log.Println("!! CONNESSIONE AL DATABASE FALLITA");
		log.Panic(err);
	}

	router := gin.Default()
	routes.SetupRoutes(router);

	router.Run(":8080")
}

