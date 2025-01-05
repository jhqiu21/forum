package users

import (
	"backend/internal/dataaccess"
	"backend/internal/models"
	"encoding/json"
	"net/http"
)

const (
	InvalidRequestBodyError = "Invalid request body"
	CreateUserError         = "Error creating user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, InvalidRequestBodyError, http.StatusBadRequest)
		return
	}
	err = users.SaveUser(user)

	if err != nil {
		http.Error(w, CreateUserError, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
