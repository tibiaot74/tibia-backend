package mappers

type OutfitInfo struct {
	Id        int
	IsPremium bool
}

var OutfitMap = map[string]OutfitInfo{
	"warrior_male": {
		Id:        134,
		IsPremium: true,
	},
	"warrior_female": {
		Id:        142,
		IsPremium: true,
	},
	"summoner_male": {
		Id:        133,
		IsPremium: true,
	},
	"summoner_female": {
		Id:        141,
		IsPremium: true,
	},
	"noble_male": {
		Id:        132,
		IsPremium: true,
	},
	"noble_female": {
		Id:        140,
		IsPremium: true,
	},
	"knight_male": {
		Id:        131,
		IsPremium: true,
	},
	"knight_female": {
		Id:        139,
		IsPremium: true,
	},
	"mage_male": {
		Id:        130,
		IsPremium: false,
	},
	"mage_female": {
		Id:        138,
		IsPremium: false,
	},
	"hunter_male": {
		Id:        129,
		IsPremium: false,
	},
	"hunter_female": {
		Id:        137,
		IsPremium: false,
	},
	"citzen_male": {
		Id:        128,
		IsPremium: false,
	},
	"citzen_female": {
		Id:        136,
		IsPremium: false,
	},
}
