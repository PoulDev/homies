package routes

import (
	"github.com/PoulDev/roommates-api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
    authRoutes.POST("/login", authLogin)
    authRoutes.POST("/register", authRegister)
    authRoutes.POST("/renew", middlewares.AuthMiddleware, authRenew)

	userRoutes := router.Group("/user")
	userRoutes.Use(middlewares.AuthMiddleware);
	userRoutes.GET("/me", userMe)
	userRoutes.GET("/house", userHouse)
	userRoutes.GET("/avatar", userAvatar)
	userRoutes.POST("/house", createHouse)
	
	cartRoutes := router.Group("/cart")
	cartRoutes.Use(middlewares.AuthMiddleware)
	cartRoutes.GET("/", myShoppingCart)

	debugRoutes := router.Group("/debug")
	debugRoutes.Use(middlewares.AuthMiddleware, middlewares.AdminMiddleware);
	debugRoutes.GET("/db/check", checkDatabase)
}
