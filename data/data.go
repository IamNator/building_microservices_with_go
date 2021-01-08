package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	CreatedOn string `json:"_"`
	UpdatedOn string `json:"_"`
	DeletedOn string `json:"_"`
}

type Products []*Product

func GetProduct() Products {
	return productList
}

// FromJson deserialize the contents of the json
func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// ToJson serialize the contents of the collection to json
func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var productList = Products{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc123",
		CreatedOn: time.Now().UTC().String(),
		DeletedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffee with milk",
		Price: 1.99,
		SKU: "ty23",
		CreatedOn: time.Now().UTC().String(),
		DeletedOn: time.Now().UTC().String(),
	},
}