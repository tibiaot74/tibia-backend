package controllers

import (
	"net/http"
	"tibia-backend/database"
	"tibia-backend/models"
	"tibia-backend/requests"

	"github.com/gin-gonic/gin"
)

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
	if err := account.HashPassword(request.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	account.Premdays = 0
	account.Lastday = 0
	account.Key = ""
	account.Warnings = 0
	account.PremiumPoints = 0
	account.BackupPoints = 0
	account.GuildPoints = 0
	account.GuildPointsStats = 0
	account.Blocked = 0
	account.GroupId = 0
	account.VipTime = 0
	account.EmailNew = ""
	account.EmailNewTime = 0
	account.EmailCode = ""
	account.NextEmail = 0
	account.Created = 0
	account.PageLastday = 0
	account.PageAccess = 0
	account.Rlname = ""
	account.Location = ""
	account.Flag = ""
	account.LastPost = 0
	account.CreateDate = 0
	account.CreateIp = 0
	account.Vote = 0

	record := database.Instance.Create(&account)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	var response requests.RegisterAccountResponse
	response.Id = account.Id
	response.Name = account.Name
	response.Email = account.Email

	context.JSON(http.StatusCreated, response)
}
