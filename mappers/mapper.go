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
