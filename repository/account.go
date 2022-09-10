package repository

import (
	"fmt"
	"tibia-backend/database"
	"tibia-backend/models"
)

func RegisterAccount(
	accountName string,
	hashedPassword string,
	email string,
) error {
	var account models.Account

	account.Name = accountName
	account.Password = hashedPassword
	account.Email = email

	// Defaults
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
		fmt.Println(record.Error)
		return record.Error
	}
	return nil
}

func GetAccount(accountName string) (*models.Account, error) {
	var account models.Account

	record := database.Instance.Where("name = ?", accountName).First(&account)
	if record.Error != nil {
		return &models.Account{}, record.Error
	}
	return &account, nil
}
