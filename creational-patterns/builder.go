package main

import "fmt"

// Produto complexo
type House struct {
	Walls   string
	Doors   string
	Windows string
}

// Builder interface
type HouseBuilder interface {
	SetWalls() HouseBuilder
	SetDoors() HouseBuilder
	SetWindows() HouseBuilder
	Build() House
}

// Concrete Builder
type ConcreteHouseBuilder struct {
	house House
}

func (b *ConcreteHouseBuilder) SetWalls() HouseBuilder {
	b.house.Walls = "Concrete Walls"
	return b
}

func (b *ConcreteHouseBuilder) SetDoors() HouseBuilder {
	b.house.Doors = "Wooden Doors"
	return b
}

func (b *ConcreteHouseBuilder) SetWindows() HouseBuilder {
	b.house.Windows = "Glass Windows"
	return b
}

func (b *ConcreteHouseBuilder) Build() House {
	return b.house
}

func mainBuilder() {
	builder := &ConcreteHouseBuilder{}
	house := builder.SetWalls().SetDoors().SetWindows().Build()
	fmt.Println(house) // {Concrete Walls Wooden Doors Glass Windows}
}
