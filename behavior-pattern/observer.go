package main

import "fmt"

// Observador
type Observer interface {
	Update(data string)
}

// Concreto Observador
type ConcreteObserver struct {
	Name string
}

func (co *ConcreteObserver) Update(data string) {
	fmt.Printf("%s recebeu: %s\n", co.Name, data)
}

// Sujeito
type Subject struct {
	observers []Observer
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Notify(data string) {
	for _, observer := range s.observers {
		observer.Update(data)
	}
}

func maiObserver() {
	subject := &Subject{}

	observer1 := &ConcreteObserver{Name: "Observer 1"}
	observer2 := &ConcreteObserver{Name: "Observer 2"}

	subject.Attach(observer1)
	subject.Attach(observer2)

	subject.Notify("Novo Evento")
}
