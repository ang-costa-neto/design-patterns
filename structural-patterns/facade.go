package main

import "fmt"

// Subsistema
type SubSystem1 struct{}

func (s SubSystem1) Operation1() string {
	return "Operation 1"
}

type SubSystem2 struct{}

func (s SubSystem2) Operation2() string {
	return "Operation 2"
}

// Facade
type Facade struct {
	system1 SubSystem1
	system2 SubSystem2
}

func (f Facade) Operation() string {
	return f.system1.Operation1() + " + " + f.system2.Operation2()
}

func mainFacade() {
	facade := Facade{}
	fmt.Println(facade.Operation()) // Operation 1 + Operation 2
}
