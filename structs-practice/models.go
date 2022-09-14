package structspractice

import (
	"fmt"
)

type Product struct {
	ID,
	Title,
	Description, Price string
}

func createProduct(title, desc, price string) *Product {
	return &Product{generateSimpleId(), title, desc, price}
}

func (product *Product) getProductInfo() {
	fmt.Printf("ID: %v\nTitle: %v\nDesc: %v\nPrice: %v$\n", product.ID, product.Title, product.Description, product.Price)
}