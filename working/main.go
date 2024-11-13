package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/flexGURU/tutorials/handlers"
)

func main() {
	add := ":8080"

	l := log.New(os.Stdout, "product api ",log.LstdFlags)

	hh := handlers.NewHello(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)

	server := http.Server{
		Addr: add,
		Handler: sm,
		ErrorLog: l,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
		IdleTimeout: 120*time.Second,
	}

	go func ()  {
		l.Printf("Server started listening at port %s", add)
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
			os.Exit(1)
		}
		
	}()

	chanSignal := make(chan os.Signal, 1)
	signal.Notify(chanSignal, os.Interrupt)
	signal.Notify(chanSignal, os.Kill)

	sig := <- chanSignal
	l.Printf("Gracefully shutdown due %s", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	server.Shutdown(tc)




}