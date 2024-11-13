package handlers

import (
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
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "error unmarshalling data", http.StatusInternalServerError)
	}


}