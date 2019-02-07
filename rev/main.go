package main

import (
	"flag"
	"fmt"
)

var aiFlag = flag.Bool("C", false, "Computer Opponent")
var allPlayedMoves = []int{0} // possible fixed array
var x = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
var y = [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
var d = [3][3]int{{1, 5, 9}, {3, 5, 7}, {}}

func main() {
	flag.Parse()
	for _, row := range x {
		fmt.Println(row)
	}
	p1, p2 := generatePlayersRandomStart()
	winner := false
	for winner == false {
		winner = p1.collectPlay()
		displayGame([]player{p1, p2})
		if winner == false {
			winner = p2.collectPlay()
			displayGame([]player{p1, p2})
		}
	}
	if p1.win {
		fmt.Println(p1.token, " Wins.") // replace w/ custom print func||atleast printf(templates)?
	} else if p2.win {
		fmt.Println(p2.token, " Wins.")
	} else {
		fmt.Println("Tie Game.")
	}
	fmt.Scanln()
}

func displayGame(players []player) {
	board := [3][3]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	for _, p := range players {
		for _, move := range p.moves {
			if move <= 3 {
				board[0][move-1] = p.token
			} else if move <= 6 {
				board[1][move-4] = p.token
			} else {
				board[2][move-7] = p.token
			}
		}
	}
	for _, row := range board {
		fmt.Println(row)
	}
}

func isNewMove(move int) bool {
	for _, idx := range allPlayedMoves {
		if move == idx {
			return false
		}
	}
	return true
}

func comparator(moveSet []int) bool { // breakdown to make most useful
	winSets := [3][3][3]int{x, y, d}
	for _, set := range winSets {
		for _, row := range set {
			matchCount := 0
			for _, val := range row {
				for _, move := range moveSet {
					if val == move {
						matchCount++
					}
				}
			}
			if matchCount == 3 {
				return true
			}
		}
	}
	return false
}
