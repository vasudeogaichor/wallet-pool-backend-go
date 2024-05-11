package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vasudeogaichor/wallet-pool-backend/internal/controllers"
	"github.com/vasudeogaichor/wallet-pool-backend/internal/services"
)

func SetupUserRoutes(router *gin.Engine) {
	userService := &services.UserService{}
	userController := &controllers.UserController{UserService: *userService}

	//Define Routes
	router.POST("/users", userController.CreateUser)
	// router.GET("/users", userController.ListUser)
}
