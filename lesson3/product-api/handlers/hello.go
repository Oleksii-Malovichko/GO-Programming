package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello requests")
	
	// read the body
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Println("[Hello] Error reading body", err)
	
		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// Write the response
	fmt.Fprintf(rw, "Hello %s\n", b)
}