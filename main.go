package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vasudeogaichor/wallet-pool-backend/internal/routes"
)

func main() {
	router := gin.Default()
	// TODO - set up similar routes for other apis
	routes.SetupUserRoutes(router)

	router.Run(":8080")
}
