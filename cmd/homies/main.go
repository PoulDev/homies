package main

import (
	"fmt"
	"log"

	"github.com/zibbadies/homies/internal/homies/config"
	"github.com/zibbadies/homies/internal/homies/db"
	"github.com/zibbadies/homies/internal/homies/routes"
	"github.com/zibbadies/homies/pkg/homies/avatar"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println(avatar.HSL2RGB(avatar.HSLColor{H: 215/360.0, S: 0.65, L: 0.80}))
	err := config.LoadConfig();
	if (err != nil) {
		log.Println("Missing environment variables!")
		log.Fatal(err)
	}

	err = db.ConnectDatabase()
	if (err != nil) {
		log.Println("Connection to the database failed!")
		log.Panic(err)
	}

	router := gin.Default()
	gin.Recovery()
	routes.SetupRoutes(router)

	log.Printf("🦜 Server listening on port %d\n ", config.HostPort)
	err = router.Run(fmt.Sprintf(":%d", config.HostPort))
	if (err != nil) {
		log.Println("Failed to start server!")
		log.Panic(err.Error())
	}
}

