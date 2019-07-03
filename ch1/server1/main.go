// Server1 is a simple webserver which answers any url with the url string
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // always call handle
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echos the Path componet of the requets url
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
