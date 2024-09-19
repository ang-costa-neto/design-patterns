package main

import "fmt"

// Implementação
type Color interface {
	ApplyColor() string
}

// Concrete implementation
type Red struct{}

func (r Red) ApplyColor() string {
	return "Vermelho"
}

// Abstração
type Shape struct {
	Color Color
}

// Alterado o nome para circleBridge para nao ocorrer erro com outro exemplo que utiliza bridge
type CircleBridge struct {
	Shape
}

func (c CircleBridge) Draw() string {
	return "Círculo " + c.Color.ApplyColor()
}

func main() {
	red := Red{}
	circle := CircleBridge{Shape{Color: red}}
	fmt.Println(circle.Draw()) // Círculo Vermelho
}
