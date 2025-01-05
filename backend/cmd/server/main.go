package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// r := router.Setup()
	fmt.Println("Listening on port 8000 at http://localhost:8000!")
	http.HandleFunc("/test", handler)
	// log.Fatalln(http.ListenAndServe(":8000", r))
}
