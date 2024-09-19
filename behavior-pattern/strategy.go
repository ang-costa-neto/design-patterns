package main

import "fmt"

// Interface Strategy
type Operation interface {
	Execute(a, b int) int
}

// Concrete Strategies
type Add struct{}

func (Add) Execute(a, b int) int {
	return a + b
}

type Multiply struct{}

func (Multiply) Execute(a, b int) int {
	return a * b
}

// Contexto
type Context struct {
	operation Operation
}

func (c *Context) SetOperation(op Operation) {
	c.operation = op
}

func (c Context) ExecuteStrategy(a, b int) int {
	return c.operation.Execute(a, b)
}

func mainStrategy() {
	context := &Context{}

	context.SetOperation(Add{})
	fmt.Println("Adição:", context.ExecuteStrategy(3, 4)) // Adição: 7

	context.SetOperation(Multiply{})
	fmt.Println("Multiplicação:", context.ExecuteStrategy(3, 4)) // Multiplicação: 12
}
