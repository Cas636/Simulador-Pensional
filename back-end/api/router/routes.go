package router

import (
	"backend/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permite todas las conexiones
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Si la solicitud es OPTIONS, responde de inmediato
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Aplicar middleware CORS primero
	r.Use(enableCORS)

	// Definir endpoints
	r.HandleFunc("/api/pension/v1/simulate", handlers.HandleSimulation).Methods("POST", "OPTIONS")

	return r
}

/*
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggerMiddleware)

	r.Handle("/api/pension/v1/simulate",
		middleware.ValidateEmployeeMiddleware(http.HandlerFunc(handlers.HandleSimulation)),
	).Methods("POST")

	return r
}
*/
