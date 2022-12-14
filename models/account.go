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

type Sex bool

const (
	Female Sex = false
	Male   Sex = true
)

type Player struct {
	Id                    int `gorm:"primarykey"`
	Name                  string
	World_id              int
	Group_id              int
	Account_id            int
	Level                 int
	Vocation              int
	Health                int
	Healthmax             int
	Experience            int
	Lookbody              int
	Lookfeet              int
	Lookhead              int
	Looklegs              int
	Looktype              int
	Lookaddons            int
	Lookmount             int
	Maglevel              int
	Mana                  int
	Manamax               int
	Manaspent             int
	Soul                  int
	Town_id               int
	Posx                  int
	Posy                  int
	Posz                  int
	Conditions            string
	Cap                   int
	Sex                   int
	Lastlogin             int
	Lastip                int
	Save                  int
	Skull                 int
	Skulltime             int
	Rank_id               int
	Guildnick             string
	Lastlogout            int
	Blessings             int
	Pvp_blessing          int
	Balance               int
	Stamina               int
	Direction             int
	Loss_experience       int
	Loss_mana             int
	Loss_skills           int
	Loss_containers       int
	Loss_items            int
	Premend               int
	Online                int
	Marriage              int
	Marrystatus           int
	Promotion             int
	Deleted               int
	Description           string
	Exphist_lastexp       int
	Exphist1              int
	Exphist2              int
	Exphist3              int
	Exphist4              int
	Exphist5              int
	Exphist6              int
	Exphist7              int
	Onlinetimetoday       int
	Onlinetime1           int
	Onlinetime2           int
	Onlinetime3           int
	Onlinetime4           int
	Onlinetime5           int
	Onlinetime6           int
	Onlinetime7           int
	Onlinetimeall         int
	Auction_balance       int
	Created               int
	Nick_verify           string
	Old_name              string
	Hide_char             int
	Comment               string
	Show_outfit           int
	Show_eq               int
	Show_bars             int
	Show_skills           int
	Show_quests           int
	Stars                 int
	Create_ip             int
	Create_date           int
	Signature             string
	Cast                  int
	CastViewers           int    `gorm:"column:CastViewers"`
	CastDescription       string `gorm:"column:CastDescription"`
	Offlinetraining_time  int
	Offlinetraining_skill int
	Broadcasting          int
	Viewers               int
	Ip                    string
}
