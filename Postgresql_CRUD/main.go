package main

import (
	"test/goP"
)

/*

*INSERT	:Add a Product
*SELECT	:Get All Products
*SELECT	:Get a Product By Id
*UPDATE	:Update a Product

 */

func main() {
	product := goP.Product{
		Id:          2,
		Title:       "Go Programming Language Book",
		Description: "It's a good book",
		Price:       36.78,
	}

	goP.InsertProduct(product)
	// goP.Hey()
}
