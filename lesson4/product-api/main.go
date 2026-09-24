import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"product-api/handlers"
)

var bindAddress = ":5555"

func main() {
	address := os.Getenv("BIND_ADDRESS")
	if address != "" {
		bindAddress = address
	}

	l := log.New(os.Stdout, "products-api", log.LstdFlags)
	// l.Println("server started")

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	sm := http.NewServeMux
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := http.Server {
		Addr: *bindAddress,
		Handler: sm,
		ErrorLog: l,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
	}

}