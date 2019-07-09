// Server3 based lissjous server
package main

import (
	"fmt"
	"image/color"
	"log"
	"net/http"
	"strconv"

	"github.com/arnoks/learngo/ch1/lissajous"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/parameter", parameterizedLissajousHandler)
	http.HandleFunc("/lissajous", lissajousHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
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

func parameterizedLissajousHandler(w http.ResponseWriter, r *http.Request) {
	var freq float64

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	f := r.FormValue("frequency")

	freq, err := strconv.ParseFloat(f, 64)

	if err != nil {
		fmt.Fprintf(w, "Invlalid frequency parameter: %s", f)
	}

	//fmt.Fprintf(w, "Frequency = %f", freq)
	var p = lissajous.Parameter{freq * 3.0, []color.Color{color.White, color.Black}}
	p.Draw(w)
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	lissajous.Lissajous(w)
}
