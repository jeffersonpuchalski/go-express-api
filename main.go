package main

import (
	"log"
	"net/http"

	"github.com/InterloperStudios/ims-api/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterUsersRoutes(r)

	log.Fatal(http.ListenAndServe(":8081", r))
}
