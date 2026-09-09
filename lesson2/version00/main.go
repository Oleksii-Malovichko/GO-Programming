package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler1(rw http.ResponseWriter, r *http.Request) {
	log.Printf("This is logging\n")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return 
	}
	fmt.Fprintf(os.Stdout, "DATA: %s\n", d)
}

func handler2(rw http.ResponseWriter, r *http.Request) {
	log.Printf("Loggin for exit\n")
	d, err := io.ReadAll(r.Body)
	if (err != nil) {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return 
	}
	fmt.Fprintf(os.Stdout, "DATA: %s\n", d)
}

func main() {
	http.HandleFunc("/", handler1)
	http.HandleFunc("/exit", handler2)
	http.ListenAndServe(":5050", nil)
}

/* func main(){
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("Hello world")
		
		d, err := io.ReadAll(r.Body) // прочитанные байты лежат в 'd'
		log.Printf("DATA: %q", d)
		log.Printf("ERROR: %v", err)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return 
		}
		fmt.Fprintf(rw, "YOUR DATA: '%s'\n", d) // отправка клиенту
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye world")
	})
	http.ListenAndServe(":5050", nil)
} */

// call the error: // curl -v -X POST -H 'Content-Length: 1' --data-binary 'hello' localhost:5050
// os.Exit(1);
/* 
rw => пришет клиенту
r.Body => читаем ОТ клиента */


/* 
как реализовать handlers для ListenAndServe (структуры и тп, serverMux)
как пользоватся этой ебучей go documentation????!!!!
*/














/* 
CAUSING ERROR!!!

вызов ошибки:
Вот это уже интереснее. Можно написать клиента, который начал отправлять body, а потом оборвал TCP-соединение.

Например, сервер ожидает:

Content-Length: 100

а клиент отправляет только несколько байт и закрывает соединение.

conn, err := net.Dial("tcp", "localhost:5050")
if err != nil {
    panic(err)
}

conn.Write([]byte(
    "POST / HTTP/1.1\r\n" +
    "Host: localhost:5050\r\n" +
    "Content-Length: 100\r\n" +
    "\r\n" +
    "hello",
))

conn.Close()
*/