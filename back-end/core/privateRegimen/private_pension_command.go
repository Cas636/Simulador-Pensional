package privateRegimen

import (
	"backend/shared/models"
	"fmt"
)

// Implementación del comando para el régimen privado
type PrivatePensionCommand struct{}

func (c *PrivatePensionCommand) Execute(e *models.Employee) {
	fmt.Println("Ejecutando simulación de pensión privada...")
	SimulatePension(e)
}
