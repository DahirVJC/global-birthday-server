package controllers

import (
	"global-birthday-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	service *services.UserService
}

func NewUsersController() *UsersController {
	usersService := services.NewUsersService()

	return &UsersController{
		service: usersService,
	}
}

func (*UsersController) GetUsers(c *gin.Context) {
	c.String(http.StatusOK, "Server is up.")
}
