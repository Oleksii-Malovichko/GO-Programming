package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodBye(l*log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g*Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye world")
	rw.Write([]byte("Byeeee\n"))
}