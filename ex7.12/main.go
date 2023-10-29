package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var html = template.Must(template.New("list").Parse(
	`
	<html>
	<body>
	<table>
		<tr>
			<th>item</th>
    		<th>price</th>
		</tr>
		{{ range $k, $v := . }}
		<tr>
			<td>{{$k}}</td>
			<td>{{$v}}</td>
		</tr>
		{{end}}
	</table>
	</body>
	</html>
`))

type dollars float32

var mux sync.Mutex

func main() {
	db := database{"shoes": 50, "socks": 50}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/update", db.updatePrice)
	http.HandleFunc("/read", db.read)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	mux.Lock()
	html.Execute(w, db)
	mux.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		mux.Lock()
		delete(db, item)
		mux.Unlock()
		w.WriteHeader(http.StatusAccepted) // 202
		fmt.Fprintf(w, "File deleted: %q\n", item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) updatePrice(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

	newPrice := req.FormValue("price")
	price, err := strconv.Atoi(newPrice)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No integer price given")
	}
	mux.Lock()
	db[item] = dollars(price)
	mux.Unlock()
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

	newPrice := req.FormValue("price")
	price, err := strconv.Atoi(newPrice)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No integer price given")
	}

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s already exists\n", item)
	}

	mux.Lock()
	db[item] = dollars(price)
	mux.Unlock()
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no item given")
	}

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s doesn't exists\n")
	}

	mux.Lock()
	fmt.Fprintf(w, "%s: %d\n", item, db[item])
}
