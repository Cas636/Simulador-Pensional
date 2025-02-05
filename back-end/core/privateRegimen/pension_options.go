package privateRegimen

import (
	constants "backend/shared/constants"
	"backend/shared/models"
	utils "backend/shared/utils"
	"math"
)

func getAverage(data *models.Simulation, e *models.Employee) (int, int) {
	weeks := ((data.Options[0].NumberOfWeeks + data.Options[1].NumberOfWeeks) / 2) - e.WeeksContributed
	return weeks, weeks / 4
}

func validateAndAddOption(data *models.Simulation, option models.Option) {
	if option.Amount >= constants.MinimumSalary {
		option.Amount = math.Round(option.Amount*100) / 100
		option.Priority = len(data.Options) + 1
		data.Options = append(data.Options, option)
	}
}

func addMinimumPensionGuarantee(weeksTotal int) models.Option {
	if weeksTotal >= constants.MinimumWeeksPrivateRegime {
		return models.Option{
			NumberOfWeeks: constants.MinimumWeeksPrivateRegime,
			Amount:        constants.MinimumSalary,
		}
	}
	return models.Option{}
}

func buildOptions(e *models.Employee, data *models.Simulation) {
	remainingWeeks := utils.CalculateRemainingWeeks(e.Age, e.PensionAge)
	remainingMonths := remainingWeeks / 4

	option := calculateMinimumAmount(e, remainingMonths)
	validateAndAddOption(data, option)

	option = models.Option{
		NumberOfWeeks: remainingWeeks + e.WeeksContributed,
		Amount:        calculateAllowanceValue(e, remainingMonths),
	}
	validateAndAddOption(data, option)

	if len(data.Options) == 2 {
		weeks, months := getAverage(data, e)

		option = models.Option{
			NumberOfWeeks: weeks + e.WeeksContributed,
			Amount:        calculateAllowanceValue(e, months),
		}
		validateAndAddOption(data, option)
	}

	option = addMinimumPensionGuarantee(e.WeeksContributed + remainingWeeks)
	validateAndAddOption(data, option)
}
