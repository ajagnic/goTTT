package main

import "fmt"

var idCounter int
var allGames []game

type game struct {
	id    int
	board [9]int
}

func test() {
	fmt.Println("test")
	for range [10]int{} {
		g := newGame()
		allGames = append(allGames, g)
	}
	fmt.Println(allGames)
}

func newGame() game {
	idCounter++
	return game{idCounter, [9]int{}}
}
