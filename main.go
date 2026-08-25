package main

import (
	"global-birthday-server/controllers"
	"global-birthday-server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Controllers
	statusController := controllers.NewStatusController()

	routes.RegisterRoutes(
		router,
		statusController,
	)

	err := router.Run("localhost:8080")

	if err != nil {
		return
	}
}
