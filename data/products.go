package data

import "time"

type Product struct {
	ID          int		`json:"id"`
	Name        string	`json:"name"`
	Description string	`json:"description"`
	Price       float32	`json:"price"`
	SKU         string	`json:"sku"`
	CreatedOn 	string  `json:"-"`
    UpdatedOn 	string  `json:"-"`
    DeletedOn 	string  `json:"-"`
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

func GetProduct() []Product {
	return productList
}

