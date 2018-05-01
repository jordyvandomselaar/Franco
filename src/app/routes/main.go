package routes

import (
	"github.com/gorilla/mux"
	"github.com/jordyvandomselaar/franco/src/app/controllers"
)

func InitRoutes(r *mux.Router) {
	hc := controllers.Home{}

	r.HandleFunc("/", hc.Index).Methods("GET")
	r.HandleFunc("/about", hc.About).Methods("GET")
}
