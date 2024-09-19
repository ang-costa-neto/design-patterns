package main

import "fmt"

// Prototipo interface
type Cloneable interface {
	Clone() Cloneable
}

// Concreto que será clonado
type Document struct {
	Content string
}

func (d *Document) Clone() Cloneable {
	clone := *d
	return &clone
}

func mainPrototype() {
	doc1 := &Document{Content: "Original"}
	doc2 := doc1.Clone().(*Document) 

	fmt.Println(doc1.Content) // Original
	fmt.Println(doc2.Content) // Original
}
