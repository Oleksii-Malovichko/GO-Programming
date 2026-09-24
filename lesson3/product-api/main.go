package main

// import "lesson3/product-api/handlers"

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"time"
)

var bindAddress = ":7777"

func main() { 
	// получить значение BIND_ADDRESS из переменной окружения
	address := os.Getenv("BIND_ADDRESS") // если задать переменную и запустить код вот так, то порт будет другой): BIND_ADDRESS=:8080 go run .
	// если BIND_ADDRESS была задана, исп ее вместо :7777
	if address != "" {
		bindAddress = address
	}

	l := log.New(os.Stdout, "products-api", log.LstdFlags) // это настройка объекта-логгера, который потом можно использовать вместе с printf
	/* 
	products-api - добавление префикса перед датой
	log.LstdFlags - это настройки формата лога (включает стандартные флаги, в часности дату и время)

	можно также завершить программу вместе с этим объектом:
	l.Fatal("something went wrong"), по смыслу (l.Println("something went wrong") os.Exit(1))
	l.Panic("smth went wrong") - запускает отдельный механизм аварийного завершения текущего выполнения (он выполнится ВСЕГДА самым последним)
	
	defer (отдельная штука в go, которая запоминает ф-ю и выполняет ее перед выходом из текущей ф-ии)
	
	*/

	// create the handlers
	ph := handlers.NewProducts(l)
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodBye(l)

	// create a new server mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gh)

	// create a new server
	s := http.Server {
		Addr: bindAddress, // configure the bind address
		Handler: sm, // set the default handler
		ErrorLog: l, // set the logger for the server
		ReadTimeout: 5 * time.Second, // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
		IdleTimeout: 120 * time.Second, // max time for the connection using TCP Keep-Alive
	}

	// start the server (отправка запуск сервера в отдельный поток?)
	go func() {
		fmt.Printf("Starting the server on port %s\n", bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // 
	l.Println("Shutting down server")
	
}

// RUNNING: BIND_ADDRESS=:8080 go run . (server) || curl localhost:8080/products | jq

// ServeHTTP implements the go http.Hanlder interface
// https://golang.org/pkg/net/http/#Handler 
// marshall: https://pkg.go.dev/encoding/json#Marshal