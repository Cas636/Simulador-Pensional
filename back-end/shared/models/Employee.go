package models

import "time"

type Employee struct {
	Salary               float64
	DateOfBirth          time.Time
	WeeksContributed     int
	Age                  int
	Gender               string
	PensionAge           int
	AccumulatedFunds     float64
	SalaryPercentageBase float64
	BasePension          float64
	HasBeneficiary       bool
	TypeBeneficiary      string
}
