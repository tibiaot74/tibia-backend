package controllers

import (
	"net/http"
	"tibia-backend/models"
	"tibia-backend/repository"
	"tibia-backend/requests"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(providedPassword *string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*providedPassword), 14)
	*providedPassword = string(hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func RegisterAccount(context *gin.Context) {
	var request requests.RegisterAccountRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var account models.Account
	account.Name = request.Name
	account.Email = request.Email
	if err := HashPassword(&request.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	err := repository.RegisterAccount(request.Name, request.Password, request.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var response requests.RegisterAccountResponse
	response.Id = account.Id
	response.Name = account.Name
	response.Email = account.Email

	context.JSON(http.StatusCreated, response)
}
