package routes

import (
	"github.com/PoulDev/homies/internal/homies/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
    authRoutes.POST("/login", authLogin)
    authRoutes.POST("/register", authRegister)
    authRoutes.POST("/renew", middlewares.AuthMiddleware, authRenew)

	userRoutes := router.Group("/user")
	userRoutes.Use(middlewares.AuthMiddleware);
	userRoutes.GET("/:id", userInfo)
	userRoutes.GET("/:id/house", userHouse)
	userRoutes.GET("/:id/overview", homeOverview)

	houseRoutes := router.Group("/house")
	houseRoutes.Use(middlewares.AuthMiddleware);
	houseRoutes.POST("/create", createHouse)
	houseRoutes.POST("/:invite", joinHouse)
	houseRoutes.GET("/:invite", inviteInfo)

	cartRoutes := router.Group("/lists")
	cartRoutes.Use(middlewares.AuthMiddleware)
	cartRoutes.GET("/", getLists)
	cartRoutes.GET("/:id/", getItems)
	cartRoutes.PUT("/:id/", newItem)

	debugRoutes := router.Group("/debug")
	debugRoutes.Use(middlewares.AuthMiddleware, middlewares.AdminMiddleware);
	debugRoutes.GET("/db/check", checkDatabase)

	datasetRoutes := router.Group("/data")
	datasetRoutes.GET("/checks-dataset", checksDataset)
}
