package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusController struct{}

func NewStatusController() *StatusController {
	return &StatusController{}
}

func (*StatusController) GetStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Server is up.")
}
