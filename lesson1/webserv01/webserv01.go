package main

import (
	"fmt"
	"io"
	"os"

	// "io"
	// "log"
	"net/http"
	// "os"
)

func handler(w http.ResponseWriter, rw *http.Request) {
	fmt.Printf("The name of the object (body): %T\n", rw.Body) // *http.body, потому что body это отдельный объект, реализующий интерфейс для чтения содержимого

	// именно поэтому нужен io.ReadAll, чтобы прочитать этот объект
	d, _ := io.ReadAll(rw.Body)
	fmt.Fprintf(os.Stdout, "Body (content): %s\n", d)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5050", nil)
}

/* func main(){
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		d, _ := io.ReadAll(r.Body) // getting data from user (body)
		log.Printf("Data: %s\n", d) // showing the data from the user in terminal (side of server)
		fmt.Fprintf(rw, "Hello %s\n", d) // sending the data back to the user
	})
	http.ListenAndServe(":5050", nil)
} */

/* Получение данных через body (со стороны клиента)
И также затем отправка этих данных + статистических данных в сторону клиента обратно

Request (запрос): curl -d "Alex?" localhost:5050

Server:
2026/09/02 14:55:46 Data: Alex?

Client (response/ответ):
Hello Alex?
*/