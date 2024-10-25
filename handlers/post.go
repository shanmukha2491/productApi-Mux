package handlers

import (
	data "gorilla/Data"
	"net/http"
)
// swagger:route POST /createProduct createProduct createNewProduct
// Creates the new product
// responses:
// 200: productResponse
func (p *Product) CreateProduct(rw http.ResponseWriter, r *http.Request) {
	ls := r.Context().Value(productKey{}).(data.ProductDetails)
	data.CreateNewProduct(ls)

}
