package server

import (
	"backend/internal/database"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartApp() {
	database.InitDB()
	defer database.CloseDB()

	r := mux.NewRouter()
	fmt.Println("Listening on port 8000 at http://localhost:8000!")
	log.Fatalln(http.ListenAndServe(":8000", r))
}
