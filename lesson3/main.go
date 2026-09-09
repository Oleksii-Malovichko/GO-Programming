package main

// import "lesson3/product-api/handlers"

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/nicholasjackson/env"
	"github.com/nicholasjackson/building-microservices-youtube/product-api/handlers"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":7777", "Bind address for the server")

func main() {
	env.Parse()

	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	// create the handlers
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	// create a new server mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// create a new server
	s := http.Server {
		Addr: *bindAddress, // configure the bind address
		Handler: sm, // set the default handler
		ErrorLog: l, // set the logger for the server
		ReadTimeout: 5 * time.Second, // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
		IdleTimeout: 120 * time.Second, // max time for the connection using TCP Keep-Alive
	}

	// start the server (отправка запуск сервера в отдельный поток?)
	go func() {
		l.Println("Starting server on port 7777")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	
}