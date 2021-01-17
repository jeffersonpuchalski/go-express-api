package routes

import (
	"net/http"

	"github.com/InterloperStudios/ims-api/repository"
	"github.com/gorilla/mux"
)

// RegisterUsersRoutes Register all user routes using a mux Router.
func RegisterUsersRoutes(r *mux.Router) {
	r.HandleFunc("/api/user", repository.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/api/user", repository.InserSingleUser).Methods(http.MethodPost)
}
