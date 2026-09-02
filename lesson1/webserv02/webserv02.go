package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d, err := io.ReadAll(r.Body) // get data from body
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		log.Printf("Data: %s\n", d) // print msg from client
		fmt.Fprintf(w, "Hello %s\n", d) // send msg to client
	})
	http.ListenAndServe(":5050", nil)
}

/* 
тут идет упоминания о случае, когда может произойти ошибка */