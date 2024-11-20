
package data

import (
	"encoding/json"
	"io"
	"time"
	"github.com/go-playground/validator/v10"

)

// swagger:model
type Product struct {
	// The id of the user
	// 
	// required: true
	
	ID          int		`json:"id"`
	Name        string	`json:"name" validate:"required"`
	Description string	`json:"description" validate:"required,min=3,max=20"`
	Price       float32	`json:"price" validate:"gt=0"`
	SKU         string	`json:"sku" validate:"required"`
	CreatedOn 	string  `json:"-"`
    UpdatedOn 	string  `json:"-"`
    DeletedOn 	string  `json:"-"`
}

type Products []*Product

func (p Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error   {
	e := json.NewDecoder(r)
	return e.Decode(p)

}

func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)

}

func AddPrdoduct(p *Product){
	p.ID = getNextID()

	productList = append(productList, p)
}

func getNextID() int  {
	lp := productList[len(productList)-1]
	return lp.ID + 1
	
}

var productList = Products{
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

func GetProduct() Products {
	return productList
}






