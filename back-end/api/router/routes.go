package router

import (
	"github.com/gorilla/mux"
	"backend/api/handlers"
	"backend/api/middleware"
	"net/http"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggerMiddleware)

	r.Handle("/api/pension/v1/simulate",
		middleware.ValidateEmployeeMiddleware(http.HandlerFunc(handlers.HandleSimulation)),
	).Methods("POST")

	return r
}
