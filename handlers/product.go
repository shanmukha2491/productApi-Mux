// Package classification of product API
//
// # Documentation for product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	data "gorilla/Data"

	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

type productKey struct{}



func (p *Product) MiddlewareValidationProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		ls := data.ProductDetails{}
		err := ls.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Cannot fetch details", http.StatusInternalServerError)
			return
		}

		err = ls.Validate()
		if err != nil {
			http.Error(
				rw,
				fmt.Sprintf("Cannot Validate details %s", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), productKey{}, ls)
		log.Println("Middleware is passed!!!")
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
