package main

import "fmt"

// Component
type Graphic interface {
	Draw() string
}

// Folha (Leaf)
type Circle struct{}

func (c Circle) Draw() string {
	return "Desenha Círculo"
}

// Composto (Composite)
type Picture struct {
	graphics []Graphic
}

func (p *Picture) Add(graphic Graphic) {
	p.graphics = append(p.graphics, graphic)
}

func (p Picture) Draw() string {
	result := "Picture contendo: "
	for _, g := range p.graphics {
		result += g.Draw() + ", "
	}
	return result
}

func mainComposite() {
	circle := Circle{}
	picture := Picture{}
	picture.Add(circle)

	fmt.Println(picture.Draw()) // Picture contendo: Desenha Círculo,
}
