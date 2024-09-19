package main

import "fmt"

// Interface componente
type Notifier interface {
	Send(message string) string
}

// Concrete Component
type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) string {
	return "Email: " + message
}

// Decorator
type SMSDecorator struct {
	Notifier Notifier
}

func (s SMSDecorator) Send(message string) string {
	return s.Notifier.Send(message) + " e SMS: " + message
}

func mainDecorator() {
	email := EmailNotifier{}
	sms := SMSDecorator{Notifier: email}

	fmt.Println(sms.Send("Hello")) // Email: Hello e SMS: Hello
}
