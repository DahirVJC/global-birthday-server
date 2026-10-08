package controllers

import (
	"fmt"
	"global-birthday-server/models"
	"global-birthday-server/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BirthdaysController struct {
	birthdaysService *services.BirthdaysService
	sessionsService  *services.SessionsService
}

func NewBirthdaysController(birthdaysService *services.BirthdaysService, sessionsService *services.SessionsService) *BirthdaysController {
	return &BirthdaysController{
		birthdaysService: birthdaysService,
		sessionsService:  sessionsService,
	}
}

func (bc *BirthdaysController) userIdFromToken(c *gin.Context) (uuid.UUID, bool) {
	accessToken := c.GetHeader("access-token")

	userId, err := bc.sessionsService.GetUser(accessToken)
	if err != nil {
		parseAndRespond(c, err)
		return uuid.Nil, false
	}

	return userId, true
}

func parseYearParam(c *gin.Context) (int, bool) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		parseAndRespond(c, fmt.Errorf("malformed year"))
		return 0, false
	}

	return year, true
}

func parseIdParam(c *gin.Context) (uuid.UUID, bool) {
	raw := c.Param("id")
	if raw == "" {
		raw = c.Param("year")
	}

	id, err := uuid.Parse(raw)
	if err != nil {
		parseAndRespond(c, fmt.Errorf("malformed id"))
		return uuid.Nil, false
	}

	return id, true
}

func (bc *BirthdaysController) GetBirthdays(c *gin.Context) {
	userId, ok := bc.userIdFromToken(c)
	if !ok {
		return
	}

	year, ok := parseYearParam(c)
	if !ok {
		return
	}

	birthdays, err := bc.birthdaysService.ListByYear(userId, year)
	if err != nil {
		parseAndRespond(c, err)
		return
	}

	c.JSON(http.StatusOK, birthdays)
}

func (bc *BirthdaysController) GetBirthday(c *gin.Context) {
	userId, ok := bc.userIdFromToken(c)
	if !ok {
		return
	}

	year, ok := parseYearParam(c)
	if !ok {
		return
	}

	id, ok := parseIdParam(c)
	if !ok {
		return
	}

	birthday, err := bc.birthdaysService.GetByYear(userId, year, id)
	if err != nil {
		parseAndRespond(c, err)
		return
	}

	c.JSON(http.StatusOK, birthday)
}

func (bc *BirthdaysController) PostBirthday(c *gin.Context) {
	userId, ok := bc.userIdFromToken(c)
	if !ok {
		return
	}

	var newBirthday models.BirthdayReq
	if err := c.ShouldBindJSON(&newBirthday); err != nil {
		parseAndRespond(c, err)
		return
	}

	if err := bc.birthdaysService.Create(userId, newBirthday); err != nil {
		parseAndRespond(c, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (bc *BirthdaysController) PutBirthday(c *gin.Context) {
	userId, ok := bc.userIdFromToken(c)
	if !ok {
		return
	}

	id, ok := parseIdParam(c)
	if !ok {
		return
	}

	var newBirthdayData models.BirthdayReq
	if err := c.ShouldBindJSON(&newBirthdayData); err != nil {
		parseAndRespond(c, err)
		return
	}

	if err := bc.birthdaysService.Update(userId, id, newBirthdayData); err != nil {
		parseAndRespond(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (bc *BirthdaysController) DeleteBirthday(c *gin.Context) {
	userId, ok := bc.userIdFromToken(c)
	if !ok {
		return
	}

	id, ok := parseIdParam(c)
	if !ok {
		return
	}

	if err := bc.birthdaysService.Delete(userId, id); err != nil {
		parseAndRespond(c, err)
		return
	}

	c.Status(http.StatusOK)
}
