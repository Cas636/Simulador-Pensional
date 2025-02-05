package publicregimen

import (
	"sync"
	"backend/shared/models"
)

var (
	pensionOptionInstance *models.Simulation
	once                  sync.Once
)

func GetPensionOptionInstance() *models.Simulation {
	once.Do(func() {
		pensionOptionInstance = &models.Simulation{}
	})
	return pensionOptionInstance
}

func resetPensionOptions() {
	if pensionOptionInstance != nil {
		pensionOptionInstance.Options = nil
	}
}
