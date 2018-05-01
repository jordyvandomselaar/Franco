package routes

import (
	"github.com/gorilla/mux"
	"github.com/jordyvandomselaar/franco/src/app/controllers"
)

// Registers the routes of this application. You can add more functions/files with
// more routes. Just make sure to call the other funcs in here.
func InitRoutes(r *mux.Router) {
	hc := controllers.Home{}

	r.HandleFunc("/", hc.Index).Methods("GET")
	r.HandleFunc("/about", hc.About).Methods("GET")
}
