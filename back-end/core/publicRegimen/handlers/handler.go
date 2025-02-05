package handlers

import "backend/shared/models"

type Handler interface {
	SetNext(handler Handler)
	Handle(e *models.Employee, data *models.Simulation)
}
