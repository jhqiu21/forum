package router

import (
	"backend/internal/handlers/posts"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/posts", posts.CreatePost).Methods("POST")
	return r
}
