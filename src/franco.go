package main

import (
	"github.com/gorilla/mux"
	"github.com/jordyvandomselaar/franco/src/app/routes"
	"github.com/jordyvandomselaar/franco/src/server"
)

func main() {
	r := mux.NewRouter()

	routes.InitRoutes(r)
	server.InitServer(r)
}
