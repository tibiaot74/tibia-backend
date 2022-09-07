package controllers

import (
	"net/http"
	"tibia-backend/auth"
	"tibia-backend/database"
	"tibia-backend/models"
	"tibia-backend/requests"

	"github.com/gin-gonic/gin"
)

func GenerateToken(context *gin.Context) {
	var request requests.GenerateTokenRequest
	var account models.Account
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// check if email exists and password is correct
	record := database.Instance.Where("email = ?", request.Email).First(&account)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	credentialError := account.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(account.Email, account.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var response requests.GenerateTokenResponse
	response.Token = tokenString

	context.JSON(http.StatusOK, response)
}
