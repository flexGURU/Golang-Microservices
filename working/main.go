package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/go-openapi/runtime/middleware"
	"github.com/flexGURU/tutorials/handlers"
	"github.com/gorilla/mux"
	gohandlers "github.com/gorilla/handlers"

)

func main() {
	add := ":8080"

	l := log.New(os.Stdout, "product api ",log.LstdFlags)

	// hh := handlers.NewHello(l)
	ph := handlers.NewProduct(l)


	// Creating the new router
	serveMux := mux.NewRouter()
	// Registeriung a GET request as a subrouter that inhertis from the parent router i.e ServeMux
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	// Handle GET requests for this specific route 
	getRouter.HandleFunc("/", ph.GetProduct)

	addRouter := serveMux.Methods(http.MethodPost).Subrouter()
	addRouter.HandleFunc("/", ph.AddPrdoduct )

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)

	getRouter.Handle("/docs", sh)

	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("../")))


	// CORS Header
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:4200"}))



	// sm := http.NewServeMux()

	// sm.Handle("/", ph)

	server := http.Server{
		Addr: add,
		Handler: corsHandler(serveMux),
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