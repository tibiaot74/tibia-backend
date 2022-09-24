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
	account.GroupId = 1
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
	return &account, record.Error
}

func GetAccountByEmail(email string) (*models.Account, error) {
	var account models.Account

	record := database.Instance.Where("email = ?", email).First(&account)
	return &account, record.Error
}

func RegisterPlayer(
	name string,
	account_id int,
	sex models.Sex,
	outfit int,
) (*models.Player, error) {
	var player models.Player

	player.Name = name
	player.World_id = 0
	player.Group_id = 1
	player.Account_id = account_id
	player.Level = 1
	player.Vocation = 0
	player.Health = 150
	player.Healthmax = 150
	player.Experience = 0
	player.Lookbody = 0
	player.Lookfeet = 0
	player.Lookhead = 0
	player.Looklegs = 0
	player.Looktype = outfit
	player.Lookaddons = 0
	player.Lookmount = 0
	player.Maglevel = 0
	player.Mana = 0
	player.Manamax = 0
	player.Manaspent = 0
	player.Soul = 100
	player.Town_id = 12
	player.Posx = 32097
	player.Posy = 32219
	player.Posz = 7
	player.Conditions = ""
	player.Cap = 400
	player.Sex = mappers.SexToInt(sex)
	player.Save = 1
	player.Skull = 0
	player.Skulltime = 0
	player.Rank_id = 0
	player.Guildnick = ""
	player.Lastlogout = 1663996367
	player.Blessings = 0
	player.Pvp_blessing = 0
	player.Balance = 0
	player.Stamina = 201660000
	player.Direction = 0
	player.Loss_experience = 100
	player.Loss_mana = 100
	player.Loss_containers = 100
	player.Loss_items = 7
	player.Premend = 0
	player.Online = 0
	player.Marriage = 0
	player.Marrystatus = 0
	player.Promotion = 0
	player.Deleted = 0
	player.Description = ""
	player.Exphist_lastexp = 0
	player.Exphist1 = 0
	player.Exphist2 = 0
	player.Exphist3 = 0
	player.Exphist4 = 0
	player.Exphist5 = 0
	player.Exphist6 = 0
	player.Exphist7 = 0
	player.Onlinetime1 = 0
	player.Onlinetime2 = 0
	player.Onlinetime3 = 0
	player.Onlinetime4 = 0
	player.Onlinetime5 = 0
	player.Onlinetime6 = 0
	player.Onlinetime7 = 0
	player.Onlinetimeall = 0
	player.Auction_balance = 0
	player.Created = int(time.Now().UTC().Unix())
	player.Nick_verify = ""
	player.Old_name = ""
	player.Hide_char = 0
	player.Comment = ""
	player.Show_outfit = 0
	player.Show_eq = 0
	player.Show_bars = 0
	player.Show_skills = 0
	player.Show_quests = 0
	player.Stars = 0
	player.Create_ip = 2147483647
	player.Create_date = 1470516668
	player.Signature = ""
	player.Cast = 0
	player.CastViewers = 0
	player.CastDescription = ""
	player.Offlinetraining_time = 0 
	player.Offlinetraining_skill = 0
	player.Broadcasting = 0
	player.Viewers = 0
	player.Ip = "0"

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
	return &player, record.Error
}

func GetPlayersInAccount(accountId int) ([]models.Player, error) {
	var players []models.Player

	records := database.Instance.Where("account_id = ?", accountId).Find(&players)
	return players, records.Error
}
