package main

import "fmt"

// Mediator interface
type Mediator interface {
	Send(message string, colleague Colleague)
}

// Colleague interface
type Colleague interface {
	Notify(message string)
}

// Concrete Colleague
type ConcreteColleague1 struct {
	mediator Mediator
}

func (c *ConcreteColleague1) Notify(message string) {
	fmt.Println("Colleague1 recebeu mensagem:", message)
}

// Concrete Colleague 2
type ConcreteColleague2 struct {
	mediator Mediator
}

func (c *ConcreteColleague2) Notify(message string) {
	fmt.Println("Colleague2 recebeu mensagem:", message)
}

// Concrete Mediator
type ConcreteMediator struct {
	colleague1 *ConcreteColleague1
	colleague2 *ConcreteColleague2
}

func (m *ConcreteMediator) Send(message string, colleague Colleague) {
	if colleague == m.colleague1 {
		m.colleague2.Notify(message)
	} else {
		m.colleague1.Notify(message)
	}
}

func mainMediator() {
	mediator := &ConcreteMediator{}

	colleague1 := &ConcreteColleague1{mediator: mediator}
	colleague2 := &ConcreteColleague2{mediator: mediator}

	mediator.colleague1 = colleague1
	mediator.colleague2 = colleague2

	colleague1.Notify("Mensagem de Teste para 1")
	colleague2.Notify("Mensagem de Teste para 2")
}
