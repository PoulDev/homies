package main

import (
	"fmt"
	"log"

	"github.com/PoulDev/homies/internal/homies/config"
	"github.com/PoulDev/homies/internal/homies/checks"
	"github.com/PoulDev/homies/internal/homies/db"
	"github.com/PoulDev/homies/internal/homies/routes"

	"github.com/gin-gonic/gin"
)

func main() {
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

	checks.Init()

	router := gin.Default()
	routes.SetupRoutes(router)

	log.Printf("🦜 Server listening on port %d\n ", config.HostPort)
	err = router.Run(fmt.Sprintf(":%d", config.HostPort))
	if (err != nil) {
		log.Println("Failed to start server!")
		log.Panic(err.Error())
	}
}

