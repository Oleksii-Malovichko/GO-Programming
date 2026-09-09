package handlers

import (
	"log"
	"net/http"
	"io"
	"fmt"
)

type Hello struct {
	l *log.Logger // по сути вызов ServerHTTP находится где-то тут (но я его не увижу)
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l};
}

// вызов gh.ServeHTTP(rw, r) я не вижу, так как он находится в стандартной библиотеке
func (h*Hello) ServeHTTP(rw http.ResponseWriter, r*http.Request){

	h.l.Println("Hello World")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return // этот return отвечает за внутреннюю ф-ю
	}
	fmt.Fprintf(rw, "Hello %s\n", d)
}

/* по сути тут ничего не изменилось, смысл в том, что автор показывает, как можно логически разделять ф-ии в отдельные файлы, чтобы код был чище 
это называется - dependency injection
*/