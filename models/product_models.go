package models

type Product struct {
	Id          int    `json:"id"`
	ProductName string `json:"product_name"`
	ProductTag  string `json:"product_tag"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
}

var ProductList []Product
var NextProductId = 1
