package data

import "testing"


// Tests to tets the json validation for adding a product
func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "mukuna",
		Description: "fewsede",
		Price: 1.5,
		SKU: "asas",
		
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}