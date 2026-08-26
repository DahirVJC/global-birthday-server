package startup

import (
	"global-birthday-server/controllers"
	"global-birthday-server/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	statusController := controllers.NewStatusController()

	routes.RegisterRoutes(
		router,
		statusController,
	)

	return router
}

func StartRouter(router *gin.Engine) {
	err := router.Run("localhost:8080")

	if err != nil {
		return
	}
}
