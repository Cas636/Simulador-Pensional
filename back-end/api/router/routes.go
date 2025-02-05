package router

import (
	"github.com/gorilla/mux"
	"backend/api/handlers"
	"backend/api/middleware"

)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggerMiddleware)

	r.HandleFunc("/api/pension/v1/simulate", handlers.HandleSimulation).Methods("POST")

	return r
}
