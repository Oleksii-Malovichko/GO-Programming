package main

import (
	"fmt"
	"log" // есть встроенная инфа о времени (логирование, например: 2026/09/07 09:30:15 some text)
	"net/http"
	"os"
)

/*  r → что клиент прислал нам
	w → что мы хотим отправить клиенту */
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\nHello client! This is a server, written on Go!!!\n\n") // sending data to the client

	log.Printf("User-Agent: %s", r.Header.Get("User-Agent")) // showing also user-agent, but via "log" lib. короче, единственное отличие - timestamp спереди

	fmt.Fprintf(os.Stdout, "Client is accepted\n") // sign that the client is come

	userAgent := r.Header.Get("User-Agent") // getting user-agent from client
	fmt.Fprintf(os.Stdout, "User-Agent: %s\n", userAgent) // showing it

	host := r.Host // getting host from the client
	fmt.Fprintf(os.Stdout, "Host: %s\n", host) // showing it
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5050", nil)
}

// тут получается используется анонимная ф-я, можно также создать ф-ю (второй аргумент) отдельно и вставить ее и будет тоже самое
/* func main(){
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request){ // logging ("Hello" will be written on the side of server)
		log.Println("Hello !!!") // Output: 2026/09/02 14:07:49 Hello
	})
	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request){
		log.Println("bye bye!!!") // Output: 2026/09/02 14:17:18 bye bye!!!
	})
	http.ListenAndServe(":5050", nil)
} */
// to check: curl.exe -v localhost:5050

/* тут показана база логирования со стороны сервера и выдача (автоматическая) ответов (200) в сторону клиента.
В зависимости от того, на какой путь (Path) будет отправлен запрос (со стороны curl), будет выбран ответ для логирования на 
стороне сервера. В данном случае: "Hello !!!" или "bye bye!!!" */