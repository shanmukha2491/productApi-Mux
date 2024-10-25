package handlers

import (
	data "gorilla/Data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 200: prodcutsResponse
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
