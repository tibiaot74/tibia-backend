package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Id       int `gorm:"primarykey"`
	Name     string
	Email    string
	Password string
	Premdays int
	Lastday  int
	Key      string
	Warnings int

	// Possibly remove
	PremiumPoints    int
	BackupPoints     int
	GuildPoints      int
	GuildPointsStats int
	Blocked          int
	GroupId          int
	VipTime          int
	EmailNew         string
	EmailNewTime     int
	EmailCode        string
	NextEmail        int
	Created          int
	PageLastday      int
	PageAccess       int
	Rlname           string
	Location         string
	Flag             string
	LastPost         int
	CreateDate       int
	CreateIp         int
	Vote             int
}

func (account *Account) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	account.Password = string(bytes)
	return nil
}
func (account *Account) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
