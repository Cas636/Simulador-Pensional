package shared

var BeneficiaryPercentageAdjustmentMale = map[string]float64{
	"Spouse":  0.15,
	"Child":   -0.08,
	"Parent":  -0.08,
	"Sibling": -0.15,
}

var BeneficiaryPercentageAdjustmentFemale = map[string]float64{
	"Spouse":  0.02,
	"Child":   -0.05,
	"Parent":  -0.05,
	"Sibling": -0.09,
}
