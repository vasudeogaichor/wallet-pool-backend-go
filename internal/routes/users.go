package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userRouter := router.Group("/users")
	{
		userRouter.POST("/", createUser)
		userRouter.GET("/:id", getUser)
		userRouter.PUT("/:id", updateUser)
		userRouter.DELETE("/:id", deleteUser)
	}
}

func createUser(c *gin.Context) {

}

func getUser(c *gin.Context) {

}

func updateUser(c *gin.Context) {

}

func deleteUser(c *gin.Context) {

}
