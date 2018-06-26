/* Server1 is a simple webserver which answers any url with the url string
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/arnoks/learngo/chap1/lissajous"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", reqDetails)
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w)
	}

	http.HandleFunc("/lissajous", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func reqDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q", k, v)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count += 1
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	count += 1
	mu.Unlock()
}
