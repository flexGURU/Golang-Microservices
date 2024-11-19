// Package classification Product API
// 
// Documentation for Product API
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
// swagger:meta
package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flexGURU/tutorials/data"
)

// A list of all products
// swagger:response productResponse
type productResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}
type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 200: productResponse

// GetProducts returns the products from the data store
func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	productList := data.GetProduct()
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "error unmarshalling data", http.StatusInternalServerError)
	}

}

// swagger
func (p *Products) AddPrdoduct(w http.ResponseWriter, r *http.Request)  {
	p.l.Println("post")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil{
		http.Error(w, "badrequtes", http.StatusBadRequest)
		return
	}

	err = prod.Validate()
	if err != nil {
		http.Error(w, fmt.Sprintf("error validating data %s", err), http.StatusBadRequest)
	}
	// data.AddPrdoduct(prod)
}