package handlers

import (
	constants "backend/shared/constants"
	"backend/shared/models"
)

type BasePercentageHandler struct {
	next Handler
}

func (h *BasePercentageHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BasePercentageHandler) Handle(e *models.Employee, data *models.Simulation) {
	percentageBase := float64(e.Salary) / float64(constants.MinimumSalary) * 0.5
	e.BasePension = (65.5 - percentageBase) / 100 // Se convierte a porcentaje

	if h.next != nil {
		h.next.Handle(e, data)
	}
}
