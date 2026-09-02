package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		d, _ := io.ReadAll(r.Body) // getting data from user (body)
		log.Printf("Data: %s\n", d) // showing the data from the user in terminal (side of server)
		fmt.Fprintf(rw, "Hello %s\n", d) // sending the data back to the user
	})
	http.ListenAndServe(":5050", nil)
}

/* Получение данных через body (со стороны клиента)
И также затем отправка этих данных + статистических данных в сторону клиента обратно

Request (запрос): curl -d "Alex?" localhost:5050

Server:
2026/09/02 14:55:46 Data: Alex?

Client (response/ответ):
Hello Alex?
*/