package publicregimen

import (
	"backend/core/publicRegimen/handlers"
	"backend/shared/models"
	utils "backend/shared/utils"
)

func SimulatePension(e *models.Employee) {
	data := GetPensionOptionInstance() 
	resetPensionOptions() 

	e.Age = utils.CalculateAge(e.DateOfBirth)

	var err error
	e.PensionAge, err = utils.GetPensionAgeByGender(e.Gender)
	if err != nil {
		return
	}

	baseHandler := &handlers.BasePercentageHandler{}
	weeksHandler := &handlers.WeeksHandler{}
	displayHandler := &handlers.DisplayHandler{}

	baseHandler.SetNext(weeksHandler)
	weeksHandler.SetNext(displayHandler)

	baseHandler.Handle(e, data)
}
