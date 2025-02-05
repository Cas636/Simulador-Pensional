package privateRegimen

import (
	constants "backend/shared/constants"
	"backend/shared/models"
	utils "backend/shared/utils"
)

var (
	monthlyProfitability = constants.AnnualProfitability / 12
	monthlyIPC           = constants.AnnualIPC / 12
	minimumSaved         = float64(constants.MinimumSalariesPrivateRegime) * float64(constants.MinimumSalary)
)

func getPercentageOfAdjust(e *models.Employee) float64 {
	var percentageAdjustment map[string]float64

	if e.HasBeneficiary == false{
		return 0.0
	}

	switch {
	case e.Gender == "M":
		percentageAdjustment = utils.BeneficiaryPercentageAdjustmentMale
	case e.Gender == "F":
		percentageAdjustment = utils.BeneficiaryPercentageAdjustmentFemale
	}

	adjustment, exists := percentageAdjustment[e.TypeBeneficiary]

	if !exists {
		adjustment = 0.0
	}
	return adjustment
}

func calculateMinimumAmount(e *models.Employee, numberOfMonths int) models.Option {
	estimatedFunds := float64(0)

	monthlyContribution := constants.MonthlyPensionContribution * e.Salary
	estimatedFunds = e.AccumulatedFunds
	for i := 0; i < numberOfMonths; i++ {
		profitability := estimatedFunds * monthlyProfitability
		devaluation := estimatedFunds * monthlyIPC
		estimatedFunds = estimatedFunds + profitability + monthlyContribution - devaluation
		if estimatedFunds >= minimumSaved {
			return models.Option{
				NumberOfWeeks: e.WeeksContributed + ((i + 1) * 4),
				Amount:        estimatedFunds * constants.RelationToMonthlyFee,
			}
		}
	}
	return models.Option{}
}

func calculateCumulatedAmount(e *models.Employee, numberOfMonths int) float64 {
	estimatedFunds := e.AccumulatedFunds
	monthlyContribution := constants.MonthlyPensionContribution * e.Salary
	for i := 0; i < numberOfMonths; i++ {
		profitability := estimatedFunds * monthlyProfitability
		devaluation := estimatedFunds * monthlyIPC
		estimatedFunds = estimatedFunds + profitability + monthlyContribution - devaluation
	}
	return estimatedFunds
}

func calculateAllowanceValue(e *models.Employee, numberOfMonths int) float64 {
	estimatedFunds := calculateCumulatedAmount(e, numberOfMonths)
	return estimatedFunds * constants.RelationToMonthlyFee * (1 + getPercentageOfAdjust(e))
}
