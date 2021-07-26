package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dan-almenar/todoapp/handlers"
)

func main() {
	l := log.New(os.Stdout, "todo-api", log.LstdFlags)
	taskHandler := handlers.NewTaskLogger(l)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", taskHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cancel()
	server.Shutdown(timeoutContext)

}
