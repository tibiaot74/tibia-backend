package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @tags    Helpers
// @summary Check if API is healthy and responding to requests
// @success 200 {string} string "Healthy!"
// @router  /health [get]
func HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Healthy!"})
}
