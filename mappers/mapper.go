package mappers

import (
	"tibia-backend/models"
)

func SexToInt(sex models.Sex) int {
	if sex {
		return 1
	}
	return 0
}

func IntToSex(sex int) models.Sex {
	if sex == 1 {
		return models.Sex(true)
	}
	return models.Sex(false)
}

var mapping = map[string]int{
	"hunter":   0,
	"mage":     1,
	"knight":   2,
	"citzen":   3,
	"nobleman": 4,
	"warrior":  5,
	"summoner": 6,
}

func StringToOutfit(outfit string) int {
	return mapping[outfit]
}

func OutfitToString(outfitAsInt int) string {
	for key, value := range mapping {
		if value == outfitAsInt {
			return key
		}
	}
	panic("The given outfit is not mapped.")
}
