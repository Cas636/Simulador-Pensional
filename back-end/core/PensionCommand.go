package core

import "backend/shared/models"

type PensionCommand interface {
	Execute(e *models.Employee)
}
