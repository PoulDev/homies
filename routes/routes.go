package routes

import (
	"github.com/PoulDev/roommates-api/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
    authRoutes.POST("/login", authLogin)
    authRoutes.POST("/register", authRegister)

	userRoutes := router.Group("/user")
	userRoutes.Use(middlewares.AuthMiddleware());
	userRoutes.GET("/me", userMe)
}
