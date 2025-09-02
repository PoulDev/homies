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
	userRoutes.POST("/:id/house", createHouse)
	//userRoutes.GET("/:id/avatar", ...)
	
	avatarRoutes := router.Group("/avatar")
	avatarRoutes.Use(middlewares.AuthMiddleware)
	avatarRoutes.GET("/:id", getAvatar)

	cartRoutes := router.Group("/lists")
	cartRoutes.Use(middlewares.AuthMiddleware)
	cartRoutes.GET("/", getLists)
	cartRoutes.GET("/:id/", getItems)
	cartRoutes.PUT("/:id/", newItem)

	debugRoutes := router.Group("/debug")
	debugRoutes.Use(middlewares.AuthMiddleware, middlewares.AdminMiddleware);
	debugRoutes.GET("/db/check", checkDatabase)
}
