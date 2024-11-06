package main

import (
	"fmt"
	"gorilla/handlers"
	"log"
	"net/http"
	"os"
	"time"

	gohandlers "github.com/gorilla/handlers"
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

	// postRouter.Use(ph.MiddlewareValidationProduct)
	getRouter.HandleFunc("/products", ph.GetProducts)
	getRouter.HandleFunc("/product/download", ph.DownloadFile)
	postRouter.HandleFunc("/product/upload", ph.UploadDocument)
	postRouter.HandleFunc("/createProduct", ph.CreateProduct)
	putRouter.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProduct)
	deleteRouter.HandleFunc("/deleteProduct/{id:[0-9]+}", ph.DeleteProduct)

	l.Println("Server Started Successfully")

	// CORS
	// ch := gohandlers.CORS(
	// 	gohandlers.AllowedOrigins([]string{"http://localhost:5000"}),
	// )(sm)

	ch := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"http://localhost:3000"}),                   // allow only this origin
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // add all the methods you need
		gohandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),           // add necessary headers
	)(sm)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      ch,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	server.ListenAndServe()

}
