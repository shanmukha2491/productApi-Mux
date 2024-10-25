package handlers

import (
	"context"
	"fmt"
	data "gorilla/Data"
	"strconv"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}
func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	// Set the response header to application/json
	rw.Header().Set("Content-Type", "application/json")

	// Fetch the products from the data package
	products := data.FetchProducts()
	productList := data.Products(products)

	// Encode the products into JSON format
	if err := productList.ToJSON(rw); err != nil {
		http.Error(rw, "Unable to marshal products to JSON", http.StatusInternalServerError)
		return
	}
}

type productKey struct{}

func (p *Product) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	ls := r.Context().Value(productKey{}).(data.ProductDetails)
	data.CreateNewProduct(ls)

}
func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	ls := data.ProductDetails{}
	err := ls.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Cannot fetch details", http.StatusInternalServerError)
		return
	}
	numId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(rw, "Input is wrong", http.StatusInternalServerError)
	}

	if err = data.UpdateProductDetails(ls, numId); err != nil {
		http.Error(rw, "Product Not found", http.StatusInternalServerError)
		return
	}

}

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
				fmt.Sprintf("Cannot Validate details %s", err) , 
				http.StatusBadRequest)
			return
		}
		
		ctx := context.WithValue(r.Context(), productKey{}, ls)
		log.Println("Middleware is passed!!!")
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
