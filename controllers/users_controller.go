package controllers

import (
	"global-birthday-server/models"
	"global-birthday-server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersService    *services.UsersService
	sessionsService *services.SessionsService
}

func NewUsersController(userService *services.UsersService, sessionService *services.SessionsService) *UsersController {
	return &UsersController{
		usersService:    userService,
		sessionsService: sessionService,
	}
}

func (uc *UsersController) GetUser(c *gin.Context) {
	accessToken := c.GetHeader("access-token")

	userId, err := uc.sessionsService.GetUser(accessToken)

	if err != nil {
		parseAndRespond(c, err)
		return
	}

	user, err := uc.usersService.GetById(userId)

	if err != nil {
		parseAndRespond(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UsersController) GetUsers(c *gin.Context) {
	users, err := uc.usersService.Get()

	if err != nil {
		parseAndRespond(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UsersController) PostUser(c *gin.Context) {
	var newUser models.UserReq

	if err := c.ShouldBindJSON(&newUser); err != nil {
		parseAndRespond(c, err)
		return
	}

	uc.usersService.Create(newUser)
}

func (uc *UsersController) PutUser(c *gin.Context) {
	accessToken := c.GetHeader("access-token")

	userId, err := uc.sessionsService.GetUser(accessToken)

	if err != nil {
		parseAndRespond(c, err)
		return
	}

	var newUserData models.UserReq

	if err := c.BindJSON(&newUserData); err != nil {
		parseAndRespond(c, err)
		return
	}

	errPut := uc.usersService.Update(userId, newUserData)

	if errPut != nil {
		parseAndRespond(c, err)
		return
	}
}

func (uc *UsersController) DeleteUser(c *gin.Context) {
	accessToken := c.GetHeader("access-token")

	userId, err := uc.sessionsService.GetUser(accessToken)

	if err != nil {
		parseAndRespond(c, err)
		return
	}

	errDel := uc.usersService.Delete(userId)

	if errDel != nil {
		parseAndRespond(c, err)
		return
	}
}
