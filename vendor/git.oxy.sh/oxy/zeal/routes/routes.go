package routes

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
}
