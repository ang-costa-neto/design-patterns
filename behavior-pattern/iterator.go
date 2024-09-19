package main

import "fmt"

// Agregado
type Collection struct {
	items []string
}

// Iterador
type Iterator struct {
	collection *Collection
	index      int
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.collection.items)
}

func (i *Iterator) Next() string {
	item := i.collection.items[i.index]
	i.index++
	return item
}

func (c *Collection) CreateIterator() *Iterator {
	return &Iterator{collection: c}
}

func mainIterator() {
	collection := &Collection{items: []string{"Item1", "Item2", "Item3"}}
	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
