package main

import "fmt"

// Interface esperada
type Target interface {
	Request() string
}

// Classe existente
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee Request"
}

// Adapter
type Adapter struct {
	Adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}

func mainAdapter() {
	adaptee := &Adaptee{}
	adapter := &Adapter{Adaptee: adaptee}
	fmt.Println(adapter.Request()) // Adaptee Request
}
