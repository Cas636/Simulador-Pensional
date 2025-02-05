package models

type Option struct {
	NumberOfWeeks int
	Priority      int
	Amount        float64
}

type Simulation struct {
	Regimen string
	Options []Option
}
