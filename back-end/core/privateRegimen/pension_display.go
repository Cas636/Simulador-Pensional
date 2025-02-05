package privateRegimen

import (
	"fmt"
	"backend/shared/models"
)

func showOptions(data *models.Simulation) {
	fmt.Println("Opciones de Pensión Generadas:")
	for _, opt := range data.Options {
		fmt.Printf("Opción: %d, Semanas: %d, Valor de la pensión: %.2f\n", opt.Priority, opt.NumberOfWeeks, opt.Amount)
	}
}
