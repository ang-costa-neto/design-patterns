package main

import "fmt"

// Product family: Button
type Button interface {
	Render() string
}

// Abstract Factory
type GUIFactory interface {
	CreateButton() Button
}

// Concrete Factory 1
type MacFactory struct{}

func (f MacFactory) CreateButton() Button {
	return MacButton{}
}

// Concrete Factory 2
type WindowsFactory struct{}

func (f WindowsFactory) CreateButton() Button {
	return WinButton{}
}

// Concrete Products
type MacButton struct{}

func (b MacButton) Render() string {
	return "Botão Mac"
}

type WinButton struct{}

func (b WinButton) Render() string {
	return "Botão Windows"
}

func mainAbstractFactory() {
	var factory GUIFactory = MacFactory{}
	button := factory.CreateButton()
	fmt.Println(button.Render()) // Botão Mac
}
