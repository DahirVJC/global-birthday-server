package routes

import (
	"global-birthday-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	statusController *controllers.StatusController) {
	api := router.Group("/api")

	RegisterStatusRoutes(api, statusController)
}
