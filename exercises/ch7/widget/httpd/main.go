package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arnoks/learngo/exercises/ch7/widget/tracks"
)

func main() {
	http.HandleFunc("/tracks", trackRequest)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func trackRequest(w http.ResponseWriter, req *http.Request) {
	sortby := req.URL.Query().Get("sort")
	tracks.SortTrackList(w, sortby)
}

func showRequest(w http.ResponseWriter, req *http.Request) {
	head := `<!DOCTYPE html>
	<html>
	  <head>
		<title>Request details</title>
	  </head>
	  <body>
	`
	foot := `  </body>
	</html>`
	fmt.Fprintln(w, head)
	fmt.Fprintf(w, "\n<h>DEBUG</h>\n<p>Method: %v Uri: %v\n</p>", req.Method, req.URL.RequestURI())
	sortby := req.URL.Query().Get("sort")
	fmt.Fprintf(w, "sort tracks by: %s <br/>", sortby)
	fmt.Fprintln(w, foot)
}
