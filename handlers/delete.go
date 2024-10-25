package handlers

import (
	data "gorilla/Data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /deleteProduct/{id} removeProduct Delete
// Removes the product based of Id
func (p *Product) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num := vars["id"]
	ls := data.Products{}

	id, err := strconv.Atoi(num)
	if err != nil {
		http.Error(rw, "Input is wrong", http.StatusInternalServerError)
		return
	}

	ls.RemoveProduct(id)
}
