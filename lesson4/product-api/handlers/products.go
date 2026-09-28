package handlers

import (
	// "encoding/json"
	"log"
	"net/http"
	"product-api/data"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return 
	}

	if r.Method == http.MethodPut {
		// expect the id in the URL
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 1 {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
	
		p.l.Println("got id", id)
	}

	// catch all (if no method is satisfied)
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p*Products) getProducts(rw http.ResponseWriter, h*http.Request) { // get method, return product list
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}


/* 
Примеры requests:
создаем новый товар через POST и он добавляется в data.Product: curl -v localhost:5555 -d '{"id":1, "name": "tea", "description":"a nice cup of tea"}' | jq
Проверить созданный объект: curl localhost:5555 | jq
(в таком состоянии сервер (точнее базу данных можно забить) одним и тем же запросом)
*/
func (p *Products) addProduct(rw http.ResponseWriter, r*http.Request) { // создать
	p.l.Println("Handle POST Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	// p.l.Printf("Prod: %#v\n", prod)
	data.AddProduct(prod)
}