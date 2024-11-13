package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/flexGURU/tutorials/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}


func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	productList := data.GetProduct()
	jsonProducts, err := json.MarshalIndent(productList, "", "  ")
	if err != nil {
		http.Error(w, "error unmarshalling data", http.StatusInternalServerError)
	}

	w.Write(jsonProducts)

}