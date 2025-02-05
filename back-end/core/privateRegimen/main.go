package privateRegimen

import (
	utils "backend/shared/utils"
	"backend/shared/models"
)

func SimulatePension(e *models.Employee) {
	e.Age = utils.CalculateAge(e.DateOfBirth)

	var err error
	e.PensionAge, err = utils.GetPensionAgeByGender(e.Gender)
	if err != nil {
		return
	}

	data := GetPensionOptionInstance()
	resetPensionOptions()

	buildOptions(e, data)

	showOptions(data)
}
