package routes

import (
	"global-birthday-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterStatusRoutes(
	router *gin.RouterGroup,
	statusController *controllers.StatusController) {
	status := router.Group("/status")
	{
		status.GET("/", statusController.GetStatus)
	}
}
