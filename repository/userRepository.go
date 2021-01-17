package repository

import (
	"encoding/json"
	"net/http"

	"github.com/InterloperStudios/ims-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllUsers Get all users from collection, and send as a json body
func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	users, err := models.ReturnAllUsers(bson.M{})

	if err != nil {
		r.Response.StatusCode = http.StatusNotFound
	}

	json.NewEncoder(w).Encode(users)

}

// InserSingleUser Inser a single user to collection.
func InserSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
