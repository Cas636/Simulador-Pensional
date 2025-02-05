package router

import (
	"backend/api/handlers"
	"net/http"
	"github.com/gorilla/mux"
	"backend/api/middleware"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggerMiddleware)
	r.Use(middleware.EnableCORS) 

	//r.HandleFunc("/api/pension/v1/simulate", handlers.HandleSimulation).Methods("POST", "OPTIONS")


	r.Handle("/api/pension/v1/simulate",
		middleware.ValidateEmployeeMiddleware(http.HandlerFunc(handlers.HandleSimulation)),
	).Methods("POST","OPTIONS")
	return r
}
