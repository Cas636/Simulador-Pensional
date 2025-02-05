package publicregimen

import (
	"backend/shared/models"
	"fmt"
)

// Implementación del comando para el régimen público
type PublicPensionCommand struct{}

func (c *PublicPensionCommand) Execute(e *models.Employee) {
	fmt.Println("\nEjecutando simulación de pensión pública...")
	SimulatePension(e)
}
