package main

import "fmt"

// Produto interface
type Product interface {
	Use() string
}

// Produto Concreto A
type ProductA struct{}

func (p ProductA) Use() string {
	return "Produto A"
}

// Produto Concreto B
type ProductB struct{}

func (p ProductB) Use() string {
	return "Produto B"
}

// Factory que retorna produtos concretos
func CreateProduct(productType string) Product {
	if productType == "A" {
		return ProductA{}
	}
	return ProductB{}
}

func mainFactory() {
	productA := CreateProduct("A")
	fmt.Println(productA.Use()) // Produto A
	productB := CreateProduct("B")
	fmt.Println(productB.Use()) // Product B
}
