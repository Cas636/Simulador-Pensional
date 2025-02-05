package main

import (
	"fmt"
	"log"
	"net/http"
	"backend/api/router"
)

func main() {
	r := router.SetupRoutes()
	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
