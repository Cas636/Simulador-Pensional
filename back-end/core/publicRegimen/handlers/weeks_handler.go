package handlers

import (
	"backend/shared/models"
	utils "backend/shared/utils"
	"math"
)

type WeeksHandler struct {
	next Handler
}

func (h *WeeksHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *WeeksHandler) Handle(e *models.Employee, data *models.Simulation) {
	weeks := [3]int{1300, 1600, 1800}
	percentagePlusAdditionalWeeks := float64(0)
	maximumWeeksPossible := e.WeeksContributed + utils.CalculateRemainingWeeks(e.Age, e.PensionAge)

	for i, w := range weeks {
		roundedValue := 0.0

		switch w {
		case 1300:
			percentagePlusAdditionalWeeks = e.BasePension
		case 1600:
			percentagePlusAdditionalWeeks = e.BasePension + 0.105
		case 1800:
			percentagePlusAdditionalWeeks = 0.8
		}

		roundedValue = math.Round(percentagePlusAdditionalWeeks*float64(e.Salary)*100) / 100

		if maximumWeeksPossible > w {
			option := models.Option{
				NumberOfWeeks: w,
				Priority:      i + 1,
				Amount:        roundedValue,
			}
			data.Options = append(data.Options, option)
		}
	}

	if h.next != nil {
		h.next.Handle(e, data)
	}
}
