package posts

import (
	"net/http"
	// "backend/internal/dataaccess"
	"backend/internal/models"
	"encoding/json"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	/**
	if err := dataaccess.InsertPost(&post); err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}
	*/
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
