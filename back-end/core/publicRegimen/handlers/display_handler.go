package handlers

import (
	"backend/shared/models"
	"fmt"
)

type DisplayHandler struct{}

func (h *DisplayHandler) SetNext(handler Handler) {}

func (h *DisplayHandler) Handle(e *models.Employee, data *models.Simulation) {
	if len(data.Options) == 0 {
		fmt.Println("No alcanzas las semanas suficientes (1300) para pensionarte en Colpensiones 🙁")
		return
	}

	fmt.Println("Opciones de Pensión Generadas:")
	for _, opt := range data.Options {
		fmt.Printf("Opción: %d, Semanas: %d, Valor de la pensión: %.2f\n", opt.Priority, opt.NumberOfWeeks, opt.Amount)
	}

}
