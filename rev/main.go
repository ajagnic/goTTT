package main

import (
	"flag"
	"fmt"
)

var aiFlag = flag.Bool("C", false, "Computer Opponent")
var allPlayedMoves = []int{0}
var x = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
var y = [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
var d = [3][3]int{{1, 5, 9}, {3, 5, 7}, {}}

func main() {
	flag.Parse()
	displayIndexMap()
	p1, p2 := generatePlayersRandomStart()
	gameOver := false
	for gameOver == false {
		//
	}
	if p1.win {
		fmt.Println(p1.token, " Wins!") // replace w/ custom print func||atleast printf(templates)?
	} else if p2.win {
		fmt.Println(p2.token, " Wins!")
	} else {
		fmt.Println("Tie Game!")
	}
	fmt.Scanln()
}

func displayIndexMap() {
	for _, row := range x {
		fmt.Println(row)
	}
}

func displayGame() {
	return // use templates for this and index map(xwinset)
}
