package main

import "fmt"

var idCounter int

type game struct {
	id    int
	board [9]int
}

func test() {
	fmt.Println("test")
	g := newGame()
	fmt.Println(g)
}

func newGame() game {
	idCounter++
	return game{idCounter, [9]int{}}
}
