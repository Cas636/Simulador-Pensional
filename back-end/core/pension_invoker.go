package core

import (
	"backend/shared/models"
	"fmt"
)

type PensionInvoker struct {
	commands []PensionCommand
}

func (i *PensionInvoker) AddCommand(command PensionCommand) {
	i.commands = append(i.commands, command)
}

func (i *PensionInvoker) ExecuteCommands(e *models.Employee) {
	for _, command := range i.commands {
		command.Execute(e)
	}
	fmt.Printf("Finalizando la simulación...\n\n")
}
