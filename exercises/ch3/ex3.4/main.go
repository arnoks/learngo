// surface webserver returning functions as svg plain
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

// Handler calls the surface function and set the appropriate content type
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w)
}
