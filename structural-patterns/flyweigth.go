package main

import "fmt"

// Flyweight interface
type ShapeFlyweigth interface {
	DrawFlyweigth(color string) string
}

// Concrete Flyweight
type CircleFlyweigth struct{}

func (c CircleFlyweigth) DrawFlyweigth(color string) string {
	return "Desenha Círculo com cor " + color
}

// Flyweight Factory
var circleFactory = make(map[string]ShapeFlyweigth)

func GetCircleFlyweigth(color string) ShapeFlyweigth {
	if _, exists := circleFactory[color]; !exists {
		circleFactory[color] = CircleFlyweigth{}
	}
	return circleFactory[color]
}

func mainFlyweigth() {
	redCircle := GetCircleFlyweigth("vermelho")
	fmt.Println(redCircle.DrawFlyweigth("vermelho")) // Desenha Círculo com cor vermelho
}
