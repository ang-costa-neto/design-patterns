package main

import "fmt"

// Comando interface
type Command interface {
	Execute()
}

// Comandos concretos
type LightOnCommand struct {
	light *Light
}

func (c LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (c LightOffCommand) Execute() {
	c.light.Off()
}

// Receptor
type Light struct{}

func (l Light) On() {
	fmt.Println("Luz ligada")
}

func (l Light) Off() {
	fmt.Println("Luz desligada")
}

// Invocador
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(c Command) {
	r.command = c
}

func (r RemoteControl) PressButton() {
	r.command.Execute()
}

func mainCommand() {
	light := &Light{}

	lightOn := LightOnCommand{light: light}
	lightOff := LightOffCommand{light: light}

	remote := RemoteControl{}

	remote.SetCommand(lightOn)
	remote.PressButton() // Luz ligada

	remote.SetCommand(lightOff)
	remote.PressButton() // Luz desligada
}
