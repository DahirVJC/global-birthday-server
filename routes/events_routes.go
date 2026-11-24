package routes

import (
	"global-birthday-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterEventsRoutes(
	router *gin.RouterGroup,
	eventsController *controllers.EventsController) {
	events := router.Group("/event")
	{
		events.GET("/", eventsController.GetEvents)
		events.GET("/:id", eventsController.GetEvent)
		events.POST("/", eventsController.PostEvent)
		events.PUT("/:id", eventsController.PutEvent)
		events.DELETE("/:id", eventsController.DeleteEvent)
	}
}
