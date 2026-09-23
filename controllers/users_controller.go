package controllers

import (
	"global-birthday-server/models"
	"global-birthday-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersService    services.UsersService
	sessionsService services.SessionsService
}

func NewUsersController(userService services.UsersService, sessionService services.SessionsService) *UsersController {
	return &UsersController{
		usersService: userService,
	}
}

func (uc *UsersController) GetUser(c *gin.Context) {
	accessToken := c.GetHeader("access-token")

	userId, err := uc.sessionsService.GetUserId(accessToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "token didn't return an user ID",
		})
		return
	}

	user, err := uc.usersService.GetById(userId)
	if err != nil {
		c.JSON(http.StatusOK, *user)
	}

}

func (uc *UsersController) GetUsers(c *gin.Context) {
	users, err := uc.usersService.Get()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UsersController) PostUser(c *gin.Context) {
	var newUser models.UserReq

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	uc.usersService.Create(newUser)
}

func (uc *UsersController) PutUser(c *gin.Context) {
	accessToken := c.GetHeader("access-token")

	userId, err := uc.sessionsService.GetUserId(accessToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "token didn't return an user ID",
		})
		return
	}

	var newUserData models.UserReq

	if err := c.BindJSON(&newUserData); err != nil {
		return
	}

	errPut := uc.usersService.Update(userId, newUserData)

	if errPut != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update user",
		})
		return
	}
}

func (uc *UsersController) DeleteUser(c *gin.Context) {
	accessToken := c.GetHeader("access-token")

	userId, err := uc.sessionsService.GetUserId(accessToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "token didn't return an user ID",
		})
		return
	}

	errDel := uc.usersService.Delete(userId)

	if errDel != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete user",
		})
		return
	}
}
