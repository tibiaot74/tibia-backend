package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tibia-backend/auth"
	"tibia-backend/repository"
	"tibia-backend/requests"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(InputPassword *string, storedPassword *string) error {
	err := bcrypt.CompareHashAndPassword([]byte(*storedPassword), []byte(*InputPassword))
	fmt.Println(*InputPassword)
	fmt.Println(*storedPassword)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GenerateToken(context *gin.Context) {
	var request requests.GenerateTokenRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	account, err := repository.GetAccount(strconv.Itoa(request.Name))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		context.Abort()
		return
	}

	if err := CheckPassword(&request.Password, &account.Password); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	tokenString, err := auth.GenerateJWT(account.Name, account.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var response requests.GenerateTokenResponse
	response.Token = tokenString

	context.JSON(http.StatusOK, response)
}
