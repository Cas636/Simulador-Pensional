package handlers

import (
	"backend/core"
	"backend/core/privateRegimen"
	"backend/core/publicRegimen"
	"backend/shared/models"
	"encoding/json"
	"net/http"
)

type SimulationResponse struct {
	Simulations []models.Simulation `json:"simulations"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func HandleSimulation(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		sendJSONError(w, "Invalid request payload (101)", http.StatusBadRequest)
		return
	}

	invoker := core.PensionInvoker{}

	invoker.AddCommand(&privateRegimen.PrivatePensionCommand{})
	invoker.AddCommand(&publicregimen.PublicPensionCommand{})

	invoker.ExecuteCommands(&employee)

	dataPublic := publicregimen.GetPensionOptionInstance()
	dataPrivate := privateRegimen.GetPensionOptionInstance()

	s1 := models.Simulation{Options: dataPrivate.Options, Regimen: "Private"}
	s2 := models.Simulation{Options: dataPublic.Options, Regimen: "Public"}

	response := SimulationResponse{Simulations: []models.Simulation{s1, s2}}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
