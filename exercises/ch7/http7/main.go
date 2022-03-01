package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/add", db.add)
	//http.HandleFunc("/delete", db.delete)
	http.HandleFunc("*", catchall)
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

func (db database) list(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w)
	mu.Lock()
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	mu.Unlock()
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

var itemList = template.Must(template.New("itemlist").Parse(`
<h1>Go Store Items</h1>
<table>
<tr style='text-aling: left'>
  <th>#</th>
  <th>Item</th>
  <th>Price</th>
</tr>
{{range .Items}}
<tr>
  <td>{{.Item}}</td>
  <td>{{.Price}}</td>
</tr>
{{end}}
</table>
`))

func (db database) htmllist(w http.ResponseWriter, req *http.Request) {

	mu.Lock()
	if err := itemList.Execute(w, db.list); err != nil {
		log.Fatal(err)
	}
	mu.Unlock()
}
