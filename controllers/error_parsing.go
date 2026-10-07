package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

func parseAndRespond(c *gin.Context, err error) {
	if err == nil {
		return
	}

	status := http.StatusInternalServerError
	msg := err.Error()

	switch {
	case strings.Contains(strings.ToLower(err.Error()), "invalid") ||
		strings.Contains(strings.ToLower(err.Error()), "not found"):
		status = http.StatusNotFound
	case strings.Contains(strings.ToLower(err.Error()), "invalid") ||
		strings.Contains(strings.ToLower(err.Error()), "malformed"):
		status = http.StatusBadRequest
	case strings.Contains(strings.ToLower(err.Error()), "unauthorized"):
		status = http.StatusUnauthorized
	case strings.Contains(strings.ToLower(err.Error()), "forbidden"):
		status = http.StatusForbidden
	}

	c.JSON(status, APIError{
		Code:    status,
		Message: msg,
	})
	c.Abort()
}
