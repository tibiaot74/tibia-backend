package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"tibia-backend/auth"
	"tibia-backend/models"
	"tibia-backend/repository"
	"tibia-backend/requests"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const MAX_PLAYERS_PER_ACCOUNT = 16

func hashPassword(providedPassword *string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*providedPassword), 14)
	*providedPassword = string(hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func hasPlayerWithName(players []models.Player, name *string) bool {
	for _, player := range players {
		if player.Name == *name {
			return true
		}
	}
	return false
}

// @tags    Account/Login
// @summary Create user account
// @param   request body     requests.RegisterAccountRequest true "Params to create account"
// @success 200     {object} requests.RegisterAccountResponse
// @router  /account [post]
func RegisterAccount(context *gin.Context) {
	var request requests.RegisterAccountRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := hashPassword(&request.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	account, err := repository.RegisterAccount(
		strconv.Itoa(*request.Name),
		request.Password,
		request.Email,
	)
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

// @tags     Account/Login
// @summary  Create player
// @Security ApiKeyAuth
// @param    request body     requests.RegisterPlayerRequest true "Params to create player"
// @success  200     {object} requests.RegisterPlayerResponse
// @router   /account/player [post]
func RegisterPlayer(context *gin.Context) {
	var request requests.RegisterPlayerRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	claims := auth.GetTokenClaims(context)
	accountName, _ := strconv.Atoi(claims.Name)

	fmt.Print(claims.ID)

	players, err := repository.GetPlayersInAccount(accountName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching players from database"})
		context.Abort()
		return
	}
	if len(players) > MAX_PLAYERS_PER_ACCOUNT {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "can't have over " + strconv.Itoa(MAX_PLAYERS_PER_ACCOUNT) + " player characters in account"})
		context.Abort()
		return
	}
	if hasPlayerWithName(players, &request.Name) {
		context.JSON(http.StatusConflict, gin.H{"error": "player name already exists"})
		context.Abort()
		return
	}

	record, err := repository.RegisterPlayer(
		request.Name,
		accountName,
		request.Sex,
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		context.Abort()
		return
	}

	var response requests.RegisterPlayerResponse
	response.Id = record.Id
	response.Name = record.Name
	response.Sex = record.Sex
	context.JSON(http.StatusCreated, response)
}
