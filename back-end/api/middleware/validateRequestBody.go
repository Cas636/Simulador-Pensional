package middleware

import (
	"backend/shared/models"
	utils "backend/shared/utils"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func ValidateEmployeeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			sendJSONError(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var employee models.Employee
		if err := json.Unmarshal(bodyBytes, &employee); err != nil {
			sendJSONError(w, "Invalid request payload (Middleware)", http.StatusBadRequest)
			return
		}

		if employee.Salary <= 0 {
			sendJSONError(w, "Salary is required and must be greater than 0", http.StatusBadRequest)
			return
		}
		if employee.DateOfBirth.IsZero() {
			sendJSONError(w, "Date of Birth is required", http.StatusBadRequest)
			return
		}
		if employee.WeeksContributed < 0 {
			sendJSONError(w, "Weeks contributed must be at least 0", http.StatusBadRequest)
			return
		}
		if employee.Gender == "" || !(employee.Gender == "M" || employee.Gender == "F"){
			sendJSONError(w, "Gender is required or invalid", http.StatusBadRequest)
		return
		}
		if employee.AccumulatedFunds < 0 {
			sendJSONError(w, "Accumulated funds must be at least 0", http.StatusBadRequest)
			return
		}
		if employee.HasBeneficiary && employee.TypeBeneficiary == "" {
			sendJSONError(w, "TypeBeneficiary is required when HasBeneficiary is true", http.StatusBadRequest)
			return
		}
		_, exists := utils.BeneficiaryPercentageAdjustmentFemale[employee.TypeBeneficiary]
		if !exists {
			sendJSONError(w, "Value of the TypeBeneficiary is invalid", http.StatusBadRequest)
		}

		next.ServeHTTP(w, r)
	})
}
