package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"version01/handlers"
)

// по сути, вся эта сложность была создана из-за безопасности. часть ответсвенности по вызовам ф-й мы передаем другим программа (библиотекам)
/* 
Программу можно по сути описать 3-я шагами:
1. Создали обработчики
        ↓
2. Сказали маршрутизатору, какой обработчик для какого URL
        ↓
3. Сказали серверу использовать этот маршрутизатор

а дальше:
client
  ↓
server
  ↓
mux
  ↓
handler
  ↓
ServeHTTP
*/
// упрощенная модель происходящего: request -> Server -> ServerMux -> Hello.ServerHTTP() / Goodbye.ServerHTTP()

func main(){
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	sm := http.NewServeMux() // создает маршрутизатор, о нем можно думать как о таблице (url: / -> handler: Hello)
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{ // теперь мы все данные для сервера (порт и тп) устанавливаем в этой структуре
		Addr: ":5050",
		Handler: sm, // то есть обратится к маршрутизатору, чтобы узнать к какому конечному хандлеру нам нужно обратиться
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// http.ListenAndServe(":5050", sm)
	go func() { // запускает сервер в отдельном go routine, так как ListenAndServe блокирующая ф-я. без go routine программа не дошла бы до обработки сигналов 
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan // ждать пока в канал прийдет сигнал
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	s.Shutdown(tc) // после получения сигнала вызываем эту ф-ю, чтобы корректно закрыть программу
}