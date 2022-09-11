package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @tags     Helpers
// @summary  Check if secured API auth is working
// @Security ApiKeyAuth
// @success  200 {string} string "pong!"
// @router   /secured/ping [get]
func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "pong"})
}
