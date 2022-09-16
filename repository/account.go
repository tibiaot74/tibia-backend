package repository

import (
	"fmt"
	"tibia-backend/database"
	"tibia-backend/mappers"
	"tibia-backend/models"
	"time"
)

func RegisterAccount(
	accountName string,
	hashedPassword string,
	email string,
) (*models.Account, error) {
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
		return &account, record.Error
	}
	return &account, nil
}

func GetAccount(accountName string) (*models.Account, error) {
	var account models.Account

	record := database.Instance.Where("name = ?", accountName).First(&account)
	if record.Error != nil {
		return &models.Account{}, record.Error
	}
	return &account, nil
}

func RegisterPlayer(
	name string,
	account_id int,
	sex models.Sex,
	outfit int,
) (*models.Player, error) {
	var player models.Player

	player.Name = name
	player.Account_id = account_id
	player.Conditions = ""
	player.Sex = mappers.SexToInt(sex)
	player.Looktype = outfit
	player.Auction_balance = 0
	player.Created = int(time.Now().UTC().Unix())
	player.Nick_verify = ""
	player.Comment = ""
	player.Signature = ""
	player.CastDescription = ""

	record := database.Instance.Create(&player)
	if record.Error != nil {
		fmt.Println(record.Error)
		return &player, record.Error
	}
	return &player, nil
}

func GetPlayer(playerName string) (*models.Player, error) {
	var player models.Player

	record := database.Instance.Where("name = ?", playerName).First(&player)
	if record.Error != nil {
		return &player, record.Error
	}
	return &player, nil
}

func ListPlayers(accountName int) (*[]models.Player, error) {
	var players []models.Player

	record := database.Instance.Where("account_id = ?", accountName).Find(&players)
	if record.Error != nil {
		fmt.Println(record.Error)
		return &players, record.Error
	}
	return &players, nil
}
