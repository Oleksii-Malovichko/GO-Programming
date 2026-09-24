package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (gh *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	gh.l.Println("Handle Goodbye requests")

	b, err := io.ReadAll(r.Body)
	if err != nil {
		gh.l.Println("[Goodbye] Error reading body", err)

		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Goodbye %s\n", b)
}