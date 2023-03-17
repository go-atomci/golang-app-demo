package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/", &greetHandler{})
	mux.Handle("/greet", &greetHandler{})
	httpServer := http.Server{
		Addr:    ":5080",
		Handler: mux,
	}

	errCh := make(chan error)
	go func() {
		log.Println("listen on :5080")
		errCh <- httpServer.ListenAndServe()
	}()

	ctx := context.Background()

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("httpServer closed")
		httpServer.Close()
	case <-quit:
		log.Println("httpserver interrupt quit, waiting 0.5s")
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal("http server shutdown: ", err)
		}
	case err := <-errCh:
		log.Fatal(err)
	}
}

type greetHandler struct{}

func (h *greetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hello, I'm sample code"))
	w.WriteHeader(http.StatusOK)
}
