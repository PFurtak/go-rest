package data

import "time"

// Product structure for items in the coffee shop
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

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Light, and frothy",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          1,
		Name:        "Espresso",
		Description: "High caffien content",
		Price:       1.99,
		SKU:         "efg321",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
