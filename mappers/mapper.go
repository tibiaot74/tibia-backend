package mappers

import (
	"errors"
	"fmt"
	"strings"
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

func SexToString(sex int) string {
	if sex == 1 {
		return "male"
	}
	return "female"
}

func GetOutfitFromSexConcat(concat string) string {
	split := strings.Split(concat, "_")
	return split[0]
}

func OutfitToString(outfit int) string {
	for outfit_sex_concat, outfit_info := range OutfitMap {
		if outfit_info.Id == outfit {
			outfitAsString := GetOutfitFromSexConcat(outfit_sex_concat)
			return outfitAsString
		}
	}
	panic("The given outfitId could not be parsed as string!")
}

func StringToOutfit(outfit string, sex int) (int, error) {
	sexAsString := SexToString(sex)
	outfit_sex_concat := fmt.Sprintf("%s_%s", strings.ToLower(outfit), sexAsString)
	outfitAsInt := OutfitMap[outfit_sex_concat].Id
	if outfitAsInt == 0 {
		return 0, errors.New("The outfit " + outfit + "does not exist!")
	}
	return outfitAsInt, nil
}
