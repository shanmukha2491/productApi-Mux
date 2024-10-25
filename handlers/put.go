package handlers

import (
	data "gorilla/Data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} updateProduct updateExistingProduct
// Creates the new product by id
// responses:
// 200: productResponse
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
