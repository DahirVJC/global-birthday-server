package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (*UsersController) GetUsers(c *gin.Context) {
	c.String(http.StatusOK, "Server is up.")
}
