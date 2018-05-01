package main

import (
	"github.com/gorilla/mux"
	"github.com/jordyvandomselaar/cms/app/routes"
	"github.com/jordyvandomselaar/cms/server"
)

func main() {
	r := mux.NewRouter()

	routes.InitRoutes(r)
	server.InitServer(r)
}
