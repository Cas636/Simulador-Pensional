package shared

import (
	consts "backend/shared/constants"
	"errors"
	"time"
)

func CalculateAge(dateOfBirth time.Time) int {
	now := time.Now()
	age := now.Year() - dateOfBirth.Year()

	if now.Month() < dateOfBirth.Month() || (now.Month() == dateOfBirth.Month() && now.Day() < dateOfBirth.Day()) {
		age--
	}
	return age

}

func CalculateRemainingWeeks(age int, pensionAge int) int {
	return (pensionAge - age) * 52
}

func GetPensionAgeByGender(gender string) (int, error) {
	switch gender {
	case consts.FemaleGender:
		return consts.PensionAgeWoman, nil
	case consts.MaleGender:
		return consts.PensionAgeMan, nil
	default:
		return 0, errors.New("género no válido")

	}
}

