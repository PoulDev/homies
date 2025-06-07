package main

import (
	"fmt"
	"log"

	"github.com/PoulDev/roommates-api/config"
	"github.com/PoulDev/roommates-api/pkg/avatar"
	"github.com/PoulDev/roommates-api/pkg/db"
	"github.com/PoulDev/roommates-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println(avatar.RandAvatar());
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
	
	err = router.Run(fmt.Sprintf(":%d", config.HostPort))
	if (err != nil) {
		log.Println(err.Error())
	}
}

