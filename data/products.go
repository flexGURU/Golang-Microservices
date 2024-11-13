package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

var productList = []Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Cold coffee",
		Price:       89.8,
		SKU:         "ABC",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Hot coffee shot",
		Price:       50.0,
		SKU:         "DEF",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}

func getProduct() []Product {
	return productList
}

