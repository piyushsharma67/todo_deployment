package routes

import (
	"todo/server"

	"github.com/gorilla/mux"
)


func Routes()*mux.Router{
	r := mux.NewRouter()

	r.HandleFunc("/health",server.HealthHandler).Methods("GET")

	return r
}