package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var instance *Singleton

// Singleton struct
type Singleton struct {
	Name string
}

// Função para obter instância do Singleton
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Name: "Única Instância"}
	})
	return instance
}

func mainSingleton() {
	singleton1 := GetInstance()
	singleton2 := GetInstance()

	fmt.Println(singleton1 == singleton2) // true
}
