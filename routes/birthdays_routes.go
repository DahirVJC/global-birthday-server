package routes

import (
	"global-birthday-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterBirthdaysRoutes(
	router *gin.RouterGroup,
	birthdaysController *controllers.BirthdaysController) {
	birthdays := router.Group("/birthday")
	{
		birthdays.GET("/:year", birthdaysController.GetBirthdays)
		birthdays.GET("/:year/:id", birthdaysController.GetBirthday)
		birthdays.POST("/", birthdaysController.PostBirthday)
		// Gin requires the same wildcard name at this path position as GET /:year.
		birthdays.PUT("/:year", birthdaysController.PutBirthday)
		birthdays.DELETE("/:year", birthdaysController.DeleteBirthday)
	}
}
