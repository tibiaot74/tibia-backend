package controllers

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"tibia-backend/auth"
	"tibia-backend/mappers"
	"tibia-backend/repository"
	"tibia-backend/requests"
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

func validatePasswordMinRequirements(providedPassword string) bool {
	hasUppercase, hasNumber, hasLowercase := false, false, false
	for _, char := range providedPassword {
		if unicode.IsNumber(char) {
			hasNumber = true
		}
		if unicode.IsLower(char) {
			hasLowercase = true
		}
		if unicode.IsUpper(char) {
			hasUppercase = true
		}
	}
	validLenght := false
	if len(providedPassword) >= 6 && len(providedPassword) <= 40 {
		validLenght = true
	}

	if hasUppercase && hasNumber && hasLowercase && validLenght {
		return true
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

	if !validatePasswordMinRequirements(request.Password) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Password should be 6 to 40 characters long and should have at least one uppercase letter, one lowercase letter, one special letter (!@#$&*) and one number!"})
		context.Abort()
		return
	}

	if err := hashPassword(&request.Password); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if _, err := repository.GetAccount(strconv.Itoa(*request.Name)); err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Account ID already exists!"})
		context.Abort()
		return
	}

	if _, err := repository.GetAccountByEmail(request.Email); err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered!"})
		context.Abort()
		return
	}

	account, err := repository.RegisterAccount(
		strconv.Itoa(*request.Name),
		request.Password,
		request.Email,
	)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error": "Account already exists!"})
		context.Abort()
		return
	}

	var response requests.RegisterAccountResponse
	response.Id = account.Id
	response.Name = account.Name
	response.Email = account.Email

	context.JSON(http.StatusCreated, response)
}

func isValidPlayerName(playerName string) bool {
	regex, _ := regexp.Compile("[A-Za-z0-9 ]")
	playerNameValidatedByRegexAsSlice := regex.FindAllString(playerName, -1)
	playerNameValidatedByRegex := strings.Join(playerNameValidatedByRegexAsSlice[:], "")

	return playerNameValidatedByRegex == playerName
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

	if len(request.Name) > 20 || len(request.Name) < 3 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Player name should have at least 3 and at max 20 characters at maximum!"})
		context.Abort()
		return
	}
	if !isValidPlayerName(request.Name) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Player name should contains no only letters without accentuation and numbers!"})
		context.Abort()
		return
	}

	outfit, err := mappers.StringToOutfit(request.Outfit, mappers.SexToInt(*request.Sex))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Provided outfit " + request.Outfit + "is not valid."})
		context.Abort()
		return
	}

	players, err := repository.GetPlayersInAccount(accountId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching players from database"})
		context.Abort()
		return
	}
	if len(players) >= MAX_PLAYERS_PER_ACCOUNT {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "can't have over " + strconv.Itoa(MAX_PLAYERS_PER_ACCOUNT) + " player characters in account"})
		context.Abort()
		return
	}

	if _, err := repository.GetPlayer(request.Name); err == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Player name " + request.Name + "already exists!"})
		context.Abort()
		return
	}

	player, err := repository.RegisterPlayer(
		request.Name,
		accountId,
		*request.Sex,
		outfit,
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
	response.Outfit = mappers.OutfitToString(player.Looktype)
	context.JSON(http.StatusCreated, response)
}

// @tags     Account/Login
// @summary  Get all players of a specific account
// @Security ApiKeyAuth
// @success  200 {object} requests.ListPlayersResponse
// @router   /account/player [get]
func ListPlayers(context *gin.Context) {
	claims := auth.GetTokenClaims(context)
	accountId := claims.Id
	var response requests.ListPlayersResponse
	players, err := repository.GetPlayersInAccount(accountId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		context.Abort()
		return
	}
	if len(players) == 0 {
		response.Players = []requests.ListPlayerInfo{}
		context.JSON(http.StatusOK, response)
		return
	}
	var playersInfo []requests.ListPlayerInfo
	for i := 0; i < len(players); i++ {
		playersInfo = append(playersInfo, requests.ListPlayerInfo{
			Name:   players[i].Name,
			Level:  players[i].Level,
			Sex:    mappers.IntToSex(players[i].Sex),
			Outfit: mappers.OutfitToString(players[i].Looktype),
		})
	}

	response.Players = playersInfo
	context.JSON(http.StatusOK, response)
}
