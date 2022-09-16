package main

import "fmt"

// !!! The bigger the interface the weaker the abstraction !!

// This only works when a type satisfies all of the methods in the interface.
type printer interface {
	print()
}

type movie struct {
	title string
	price int
}

func (m *movie) print() {
	fmt.Println(m.title)
}

type game struct {
	title string
	price int
}

func (g *game) print() {
	fmt.Println(g.title)
}

func (g *game) talk() {
	fmt.Println(g.title)
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry we are still waiting")
		return
	}

	for _, it := range l {
		it.print()
	}
}

func main() {
	var (
		broxTail  = movie{title: "broxTail", price: 50}
		minecraft = game{title: "minecraft", price: 10}
		tetris    = game{title: "tetris", price: 20}
	)

	var store list

	store = append(store, &minecraft, &tetris, &broxTail)
	store.print()

}
