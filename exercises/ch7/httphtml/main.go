package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/add", db.add)
	//	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("*", catchall)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func catchall(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "\nDEBUG\nMethod: %v Uri: %v\n", req.Method, req.URL.RequestURI())
}

var mu sync.Mutex

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var itemList = template.Must(template.New("itemList").Parse(`
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{ .Titel }}</title>
		<link rel="stylesheet" href="/css/style.css">
	</head>
	<body>
		<h1> {{ .Titel }}</h1>
		<table>
		<tr><th>Item<th/><th>Price<th/></tr>
		{{range $k,$v := .List}}
		    <tr><td>{{ $k }}<td/><td>{{ $v }}</td></tr>
		{{end}}
		</table>
	</body>
</html>`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	data := struct {
		Titel string
		List  map[string]dollars
	}{
		"Golang Pricelist Example",
		db,
	}
	mu.Lock()
	err := itemList.Execute(w, data)
	mu.Unlock()
	if err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("Item")
	mu.Lock()
	price, ok := db[item]
	mu.Unlock()

	fmt.Fprintln(w)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) add(w http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "\nDEBUG\nMethod: %v Uri: %v\n", req.Method, req.URL.RequestURI())
		return
	}

	item := req.PostFormValue("Item")

	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "no item specified")
		return
	}

	sprice := req.PostFormValue("Price")
	if sprice == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "no price specified")
		return
	}

	price, err := strconv.ParseFloat(sprice, 32)
	if err != nil {
		fmt.Fprint(w, "invalid price")
		return
	}
	mu.Lock()
	db[item] = dollars(price)
	mu.Unlock()
}
