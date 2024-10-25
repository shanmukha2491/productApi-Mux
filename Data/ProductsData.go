package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
)

type ProductDetails struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"  validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	CreateOn    string  `json:"createdOn"`
	UpdatedOn   string  `json:"updatedOn"`
	DeletedOn   string  `json:"deletedOn"`
}

func (p *ProductDetails) Validate() error {

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(p)
	return err
}

type Products []ProductDetails

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *ProductDetails) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func FetchProducts() []ProductDetails {
	return products
}

func CreateNewProduct(p ProductDetails) {
	p.Id = fetchLastId()
	products = append(products, p)
}

var ErrorProductNotFound = fmt.Errorf("Product Not Found")

func UpdateProductDetails(p ProductDetails, id int) error {

	for i, prod := range products {
		if prod.Id == id {
			p.Id = id
			products[i] = p
			return nil
		}

	}

	return ErrorProductNotFound

}

func fetchLastId() int {
	id := products[len(products)-1].Id
	return id + 1
}

var products = []ProductDetails{
	{
		Id:          1,
		Name:        "Wireless Mouse",
		Description: "Ergonomic wireless mouse with a comfortable grip.",
		Price:       29.99,
		CreateOn:    time.Now().Format("2006-01-02 15:04:05"),
		UpdatedOn:   time.Now().Format("2006-01-02 15:04:05"),
		DeletedOn:   "",
	},
	{
		Id:          2,
		Name:        "Bluetooth Headphones",
		Description: "Noise-cancelling over-ear headphones with great sound quality.",
		Price:       89.99,
		CreateOn:    time.Now().Format("2006-01-02 15:04:05"),
		UpdatedOn:   time.Now().Format("2006-01-02 15:04:05"),
		DeletedOn:   "",
	},
	{
		Id:          3,
		Name:        "HD Monitor",
		Description: "27-inch Full HD monitor with ultra-slim design.",
		Price:       199.99,
		CreateOn:    time.Now().Format("2006-01-02 15:04:05"),
		UpdatedOn:   time.Now().Format("2006-01-02 15:04:05"),
		DeletedOn:   "",
	},
}
