package main

import "fmt"

// Interface State
type State interface {
	Handle() string
}

// Concrete States
type OnState struct{}

func (s OnState) Handle() string {
	return "Ligado"
}

type OffState struct{}

func (s OffState) Handle() string {
	return "Desligado"
}

// Context
type ContextState struct {
	state State
}

func (c *ContextState) SetState(state State) {
	c.state = state
}

func (c ContextState) Request() string {
	return c.state.Handle()
}

func mainState() {
	context := &ContextState{}

	context.SetState(OnState{})
	fmt.Println(context.Request()) // Ligado

	context.SetState(OffState{})
	fmt.Println(context.Request()) // Desligado
}
