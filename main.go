package main

import (
	"fmt"
	"gorilla/handlers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to gorillaMux")
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProduct(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	postRouter := sm.Methods(http.MethodPost).Subrouter()

	postRouter.Use(ph.MiddlewareValidationProduct)

	getRouter.HandleFunc("/products", ph.GetProducts)
	postRouter.HandleFunc("/createProduct", ph.CreateProduct)
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProduct)
	deleteRouter.HandleFunc("/deleteProduct/{id:[0-9]+}", ph.DeleteProduct)

	l.Println("Server Started Successfully")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	server.ListenAndServe()

}
