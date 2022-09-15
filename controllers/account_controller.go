package controllers

import (
	"net/http"
	"strconv"
	"tibia-backend/auth"
	"tibia-backend/mappers"
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

	if err := HashPassword(&request.Password); err != nil {
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
	accountId := claims.Id

	_, err := repository.GetPlayer(request.Name)
	if err == nil {
		context.JSON(http.StatusConflict, gin.H{"error": "player name already exists"})
		context.Abort()
		return
	}

	player, err := repository.RegisterPlayer(
		request.Name,
		accountId,
		*request.Sex,
		mappers.StringToOutfit(request.Outfit),
	)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		context.Abort()
		return
	}

	var response requests.RegisterPlayerResponse
	response.Id = player.Id
	response.Name = player.Name
	response.Sex = mappers.IntToSex(player.Sex)
	response.Outfit = mappers.OutfitToString(player.Lookbody)
	context.JSON(http.StatusCreated, response)
}

// @tags     Account/Login
// @summary  Get all players of a specific account
// @Security ApiKeyAuth
// @success  200     {object} requests.ListPlayersResponse
// @router   /account/player [get]
func ListPlayers(context *gin.Context) {
	claims := auth.GetTokenClaims(context)
	accountId := claims.Id
	var response requests.ListPlayersResponse

	players, err := repository.ListPlayers(accountId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		context.Abort()
		return
	}
	if len(*players) == 0 {
		response.Players = []requests.ListPlayerInfo{}
		context.JSON(http.StatusOK, response)
		return
	}

	var playersInfo []requests.ListPlayerInfo
	for i := 0; i < len(*players); i++ {
		playersInfo = append(playersInfo, requests.ListPlayerInfo{
			Name:   (*players)[i].Name,
			Level:  (*players)[i].Level,
			Sex:    mappers.IntToSex((*players)[i].Sex),
			Outfit: mappers.OutfitToString((*players)[i].Lookbody),
		})
	}

	response.Players = playersInfo
	context.JSON(http.StatusOK, response)
}
