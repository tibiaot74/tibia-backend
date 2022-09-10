package models

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
