package main

import "fmt"

// Template
type Game interface {
	Init()
	Start()
	End()
}

// Classe abstrata
type GameTemplate struct{}

func (g GameTemplate) Play(game Game) {
	game.Init()
	game.Start()
	game.End()
}

// Concrete Game
type Football struct {
	GameTemplate
}

func (f Football) Init() {
	fmt.Println("Iniciando futebol")
}

func (f Football) Start() {
	fmt.Println("Jogando futebol")
}

func (f Football) End() {
	fmt.Println("Encerrando futebol")
}

func mainTemplateMethod() {
	football := Football{}
	football.Play(football) // Executa a sequência: Iniciar -> Jogar -> Encerrar
}
