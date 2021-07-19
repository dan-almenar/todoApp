package main

import (
	"log"
	"os"
	"os/signal"
	"net/http"
	"github.com/dan-almenar/todoapp/handlers"
	"time"
	"context"
)

func main(){
	l := log.New(os.Stdout, "todo-api", log.LstdFlags)
	taskHandler := handlers.NewTaskLogger(l)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", taskHandler)

	server := &http.Server{
		Addr: ":8080",
		Handler: serveMux,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func(){
		err := server.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()
	
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)	
}