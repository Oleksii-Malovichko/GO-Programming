package main

import "net/http"
import "log"

func main(){
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request){ // logging ("Hello" will be written on the side of server)
		log.Println("Hello !!!") // Output: 2026/09/02 14:07:49 Hello
	})
	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request){
		log.Println("bye bye!!!") // Output: 2026/09/02 14:17:18 bye bye!!!
	})
	http.ListenAndServe(":5050", nil)
}
// to check: curl.exe -v localhost:5050

/* тут показана база логирования со стороны сервера и выдача (автоматическая) ответов (200) в сторону клиента.
В зависимости от того, на какой путь (Path) будет отправлен запрос (со стороны curl), будет выбран ответ для логирования на 
стороне сервера. В данном случае: "Hello !!!" или "bye bye!!!" */